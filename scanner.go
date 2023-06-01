package main

import (
	"fmt"
)

type Stack struct {
	n   int
	isi []string
}

func push(stack *Stack, x string) {
	stack.n++
	stack.isi = append(stack.isi, x)
}

func pop(stack *Stack) string {
	x := stack.isi[stack.n]
	stack.n--
	stack.isi = stack.isi[:stack.n+1]
	return x
}

func isEmpty(stack Stack) bool {
	return stack.n == 0
}

func readTop(stack Stack) string {
	return stack.isi[stack.n]
}

func getToken(teks string, j *int) string {
	k := len(teks)
	simbol := string(teks[*j])
	*j++
	return simbol
}

func main() {
	tabelParsing := map[string]map[string]string{
		"S": {
			"x": "AB",
			"y": "AB",
			"a": "CD",
			"b": "CD",
			"#": "-",
		},
		"A": {
			"x": "xA",
			"y": "y",
			"a": "-",
			"b": "-",
			"#": "-",
		},
		"B": {
			"x": "x",
			"y": "yB",
			"a": "-",
			"b": "-",
			"#": "-",
		},
		"C": {
			"x": "-",
			"y": "-",
			"a": "aD",
			"b": "b",
			"#": "-",
		},
		"D": {
			"x": "-",
			"y": "-",
			"a": "a",
			"b": "bC",
			"#": "-",
		},
	}
	simbolAwal := "S"

	input := "xyyx"
	i := 0
	hsl := ""
	stack := Stack{}
	push(&stack, simbolAwal)
	simbol := getToken(input, &i)
	for !isEmpty(stack) {
		top := readTop(stack)
		if top >= "a" { // top = terminal
			if top == simbol {
				pop(&stack)
				simbol = getToken(input, &i)
			} else {
				hsl = "Error/Ditolak"
				break
			}
		} else if top <= "Z" { // top = non terminal
			sel := tabelParsing[top][simbol]
			if sel != "-" {
				pop(&stack)
				for k := len(sel) - 1; k >= 0; k-- {
					push(&stack, string(sel[k]))
				}
			} else {
				hsl = "Error/Ditolak"
				break
			}
		}
	}

	fmt.Println("<script>")
	fmt.Println("document.getElementById('hasil').innerHTML='';")
	fmt.Println("</script>")
	fmt.Println("push(&stack,simbolAwal)")
	fmt.Println("simbol=getToken(input,&i)")
	fmt.Println("for !isEmpty(stack) {")
	fmt.Println("top=readTop(stack)")
	fmt.Println("if top>='a' {")
	fmt.Println("if top==simbol {")
	fmt.Println("pop(&stack)")
	fmt.Println("simbol=getToken(input,&i)")
	fmt.Println("} else {")
	fmt.Println("hsl='Error/Ditolak'")
	fmt.Println("break")
	fmt.Println("}")
	fmt.Println("} else if top<='Z' {")
	fmt.Println("sel=tabelParsing[top][simbol]")
	fmt.Println("if sel!='-' {")
	fmt.Println("pop(&stack)")
	fmt.Println("for k:=len(sel)-1;k>=0;k-- {")
	fmt.Println("push(&stack,string(sel[k]))")
	fmt.Println("}")
	fmt.Println("} else {")
	fmt.Println("hsl='Error/Ditolak'")
	fmt.Println("break")
	fmt.Println("}")
	fmt.Println("}")
	fmt.Println("}")

	fmt.Println("<script>")
	fmt.Println("document.getElementById('tes').innerHTML='" + input[:len(input)-1] + "';")
	if simbol == "#" && hsl == "" {
		fmt.Println("document.getElementById('hasil').innerHTML='DITERIMA';")
	} else {
		fmt.Println("document.getElementById('hasil').innerHTML='" + hsl + "';")
	}
	fmt.Println("</script>")
}
