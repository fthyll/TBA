package main

import "fmt"

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

func PrintTable(tbl map[string]map[string]string) {
	fmt.Println("<table border='1'>")
	for jdlBaris, dt := range tbl {
		fmt.Println("<tr><th></th>")
		for jdl := range dt {
			if jdl == "#" {
				fmt.Println("<th>EOS</th>")
			} else {
				fmt.Println("<th>" + jdl + "</th>")
			}
		}
		fmt.Println("</tr>")
		break
	}
	for jdlBaris, dt := range tbl {
		fmt.Println("<tr><td>" + jdlBaris + "</td>")
		for jdl, isi := range dt {
			fmt.Println("<td>" + isi + "</td>")
		}
		fmt.Println("</tr>")
	}
	fmt.Println("</table>")
}

func ParseInput(input string, tabelParsing map[string]map[string]string, simbolAwal string) string {
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

	if simbol == "#" && hsl == "" {
		return "DITERIMA"
	} else {
		return hsl
	}
}
