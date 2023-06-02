package main

import (
	"github.com/gin-gonic/gin"
)

type Stack struct {
	n   int
	isi []interface{}
}

func (stack *Stack) push(x interface{}) {
	stack.n++
	stack.isi = append(stack.isi, x)
}

func (stack *Stack) pop() interface{} {
	x := stack.isi[stack.n]
	stack.n--
	return x
}

func (stack *Stack) isEmpty() bool {
	return stack.n == 0
}

func (stack *Stack) readTop() interface{} {
	return stack.isi[stack.n]
}

func getToken(text string, j *int) string {
	k := len(text)
	word := ""

	for *j < k && (text[*j] == ' ' || text[*j] == '\r' || text[*j] == '\n') {
		*j++
	}

	for *j < k && text[*j] != ' ' && text[*j] != '\r' && text[*j] != '\n' {
		word += string(text[*j])
		*j++
	}

	return word
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	type Result struct {
		Teks  string
		Hasil string
	}

	r.POST("/parse", func(c *gin.Context) {
		inputText := c.PostForm("teks")

		stack := &Stack{
			n:   0,
			isi: make([]interface{}, 0),
		}

		push := func(x interface{}) {
			stack.push(x)
		}

		pop := func() interface{} {
			return stack.pop()
		}

		isEmpty := func() bool {
			return stack.isEmpty()
		}

		readTop := func() interface{} {
			return stack.readTop()
		}

		i := 0
		hsl := ""

		push("S")
		simbol := getToken(inputText, &i)

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

		for !isEmpty() {
			top := readTop()
			if top.(string) >= "a" {
				if top.(string) == simbol {
					pop()
					simbol = getToken(inputText, &i)
				} else {
					hsl = "Error/Ditolak"
					break
				}
			} else if top.(string) <= "Z" {
				production := tabelParsing[top.(string)][simbol]
				if production != "-" {
					pop()
					for k := len(production) - 1; k >= 0; k-- {
						push(string(production[k]))
					}
				} else {
					hsl = "Error/Ditolak"
					break
				}
			}
		}

		if simbol == "#" && hsl == "" {
			hsl = "DITERIMA"
		}

		result := Result{
			Teks:  inputText,
			Hasil: hsl,
		}

		c.HTML(200, "result.html", result)
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.Run()
}
