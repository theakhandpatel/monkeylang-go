package repl

import (
	"bufio"
	"fmt"

	"io"

	"github.com/theakhandpatel/interpreter/evaluator"
	"github.com/theakhandpatel/interpreter/lexer"
	"github.com/theakhandpatel/interpreter/object"
	"github.com/theakhandpatel/interpreter/parser"
)

const ANGRY_FACE = `【≽ܫ≼】`

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, ANGRY_FACE)
	io.WriteString(out, "Woops! We ran into some funky business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
