package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func getTokenFromString(text string, index int) string {
	k := len(text)
	word := ""

	// Ignore spaces, carriage returns, line breaks, and tabs
	for index < k && (text[index] == ' ' || text[index] == '\r' || text[index] == '\n' || text[index] == '\t') {
		index++
	}

	// Extract a token
	for index < k && text[index] != ' ' && text[index] != '\r' && text[index] != '\n' && text[index] != '\t' {
		if text[index] == '>' || text[index] == '<' || text[index] == '=' || text[index] == '+' || text[index] == '-' {
			if word != "" {
				return word
			} else {
				word += string(text[index])
				index++
				return word
			}
		}
		word += string(text[index])
		index++
	}

	return word
}

func scan(text string) []string {
	var tokens []string
	k := len(text)
	j := 0

	for j < k {
		token := getTokenFromString(text, j)
		tokens = append(tokens, token)
		j += len(token)
	}

	return tokens
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("index.html", "result.html")

	r.POST("/parse", func(c *gin.Context) {
		inputText := c.PostForm("teks")
		tokens := scan(inputText)
		fmt.Println(tokens) // Untuk melihat tokens yang dihasilkan
		c.HTML(200, "result.html", gin.H{
			"tokens": strings.Join(tokens, "\n"),
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/result", func(c *gin.Context) {
		tokens := c.Query("tokens")
		c.HTML(200, "result.html", gin.H{
			"tokens": tokens,
		})
	})

	r.Run()
}
