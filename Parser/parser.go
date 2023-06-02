package main

import (
	"github.com/gin-gonic/gin"
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
	stack.n--
	x := stack.isi[stack.n]
	stack.isi = stack.isi[:stack.n]
	return x
}

func isEmpty(stack *Stack) bool {
	return stack.n == 0
}

func readTop(stack *Stack) string {
	return stack.isi[stack.n-1]
}

func getToken(teks string, j *int) string {
	simbol := string(teks[*j])
	*j++
	return simbol
}

func parseInput(teks string) string {
	stack := Stack{
		n:   0,
		isi: []string{},
	}

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
	input := teks + "#"

	push(&stack, simbolAwal)
	i := 0
	simbol := getToken(input, &i)

	for !isEmpty(&stack) {
		top := readTop(&stack)

		if top >= "a" {
			if top == simbol {
				pop(&stack)
				simbol = getToken(input, &i)
			} else {
				return "Error/Ditolak"
			}
		} else if top <= "Z" {
			sel := tabelParsing[top][simbol]
			if sel != "-" {
				pop(&stack)
				for k := len(sel) - 1; k >= 0; k-- {
					push(&stack, string(sel[k]))
				}
			} else {
				return "Error/Ditolak"
			}
		}
	}

	if simbol == "#" {
		return "DITERIMA"
	}

	return "Error/Ditolak"
}

func main() {
	html := `<!DOCTYPE html>
	<html>
		<head>
			<title>LL(1) Parser</title>
			<style>
				td {
					width: 50px;
					text-align: center;
				}
			</style>
		</head>
		<body>
			<form method='post'>
				Aturan Produksi : <br>
				<textarea name='cfg' rows="7" cols="34" readonly>
LL(1) DENGAN NOTASI SEDERHANA
S -> AB | CD
A -> xA | y
B -> yB | x  
C -> aD | b  
D -> bC | a
				</textarea><br><br>
				Parse Table<br>
				<!-- Display the parse table here -->
				<br>
				Contoh input diterima:<br>
				yx, xyx, xyyx, xxyyx<br><br>
				Input : <br>
				<textarea name='teks' id='tes' rows="3"></textarea><br><br>
				<input type='submit' value='Parsing'><br>
				Hasil : <br>
				<textarea id='hasil' rows="3" readonly></textarea>
			</form>
		
			<script>
				function updateOutput(text) {
					document.getElementById('hasil').value = text;
				}
		
				function parseInput() {
					var input = document.getElementById('tes').value;
					updateOutput('');
		
					fetch('/', {
						method: 'POST',
						headers: {
							'Content-Type': 'application/x-www-form-urlencoded',
						},
						body: 'teks=' + encodeURIComponent(input),
					})
						.then(function (response) {
							return response.text();
						})
						.then(function (text) {
							updateOutput(text);
						})
						.catch(function (error) {
							console.error('Error:', error);
						});
				}
		
				document.querySelector('form').addEventListener('submit', function (e) {
					e.preventDefault();
					parseInput();
				});
			</script>
		</body>
	</html>`

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(200, html)
	})

	router.POST("/", func(c *gin.Context) {
		teks := c.PostForm("teks")
		result := parseInput(teks)
		c.String(200, result)
	})

	router.Run(":8080")
}
