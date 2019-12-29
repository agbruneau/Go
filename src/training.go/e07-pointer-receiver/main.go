package main

import (
	"fmt"
	"strings"
)

type Post struct {
	Title     string
	Text      string
	published bool
}

func (p Post) Headline() string {
	return fmt.Sprintf("%v - %v", p.Title, p.Text[:50])
}

func (p *Post) Published() bool { return p.published }

func (p *Post) Publish() {
	p.published = true
}

func (p *Post) Unpublish() {
	p.published = true
}

func UpperTitle(p *Post) {
	p.Title = strings.ToUpper(p.Title)
}

func main() {
	p := Post{
		Title: "Go release",
		Text: `Go is a programming language designed by Google engineers Robert Griesemer, Rob Pike, and Ken Thompson.[11] 
Go is statically typed, compiled, and syntactically similar to C, with the added benefits of memory safety, garbage collection, structural typing,[5] and CSP-style concurrency.[15] The compiler, tools, and source code are all free and open source.[16		
`,
	}

	fmt.Println(p.Headline())

	fmt.Printf("Post published? %v\n", p.Published())
	p.Publish()
	fmt.Printf("Post published? %v\n", p.Published())

	pythonPost := &Post{
		Title: "Python Intro",
		Text: `Python is an interpreted high-level programming language for general-purpose programming. Created by Guido van Rossum and first released in 1991, Python has a design philosophy that emphasizes code readability, notably using significant whitespace. It provides constructs that enable clear programming on both small and large scales.[27] In July 2018, Van Rossum stepped down as the leader in the language community after 30 years
		`,
	}

	UpperTitle(pythonPost)
	fmt.Println(pythonPost.Headline())
}
