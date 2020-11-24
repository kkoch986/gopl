package app

// Define the interactive shell used for querying

import (
	//	"errors"
	"fmt"
	"log"

	"github.com/c-bata/go-prompt"

	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/lexer"
	"github.com/kkoch986/gopl/parser"
	"github.com/kkoch986/gopl/resolver"
)

type QueryCLI struct{}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		//     {Text: "users", Description: "Store the username and age"},
		//   {Text: "articles", Description: "Store the article text posted by user"},
		// {Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}


func execCommand(t string) {
    // lex and parse the input
    l := lexer.New([]rune("?- " + t))
    if bsrSet, errs := parser.Parse(l); len(errs) > 0 {
        log.Println(errs)
        //return errors.New("Parser Error")
    } else {
        a := ast.BuildStatementList(bsrSet.GetRoot())
        output := make(chan *resolver.Bindings, 1)
        r := resolver.New()
        log.Println("Resolving...")
        go r.ResolveStatementList(a, &resolver.Bindings{}, output)
        for v := range output {
            fmt.Println("HERE", v)
            t := prompt.Input(">", completer)
            if t != ";" {
                return
            }
        }
        fmt.Println("No.")
    }
}

func exitChecker(t string, breakline bool) bool {
    if !breakline {
        return t == "quit."
    }
    return false
}

func (q *QueryCLI) Run() error {
	log.Println("Welcome to GoPL")

    qPrompt := prompt.New(
        execCommand,
        completer,
        prompt.OptionPrefix("?- "),
        prompt.OptionSetExitCheckerOnInput(exitChecker),
    )
	// Enter a REPL
    qPrompt.Run()

	return nil
}
