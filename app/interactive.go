package app

// Define the interactive shell used for querying

import (
	//	"errors"
	"fmt"
	"log"

	"github.com/c-bata/go-prompt"

	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/indexer"
	"github.com/kkoch986/gopl/lexer"
	"github.com/kkoch986/gopl/parser"
	"github.com/kkoch986/gopl/resolver"
)

type QueryCLI struct {
	I indexer.Indexer
	R *resolver.R
	H *history
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		//     {Text: "users", Description: "Store the username and age"},
		//   {Text: "articles", Description: "Store the article text posted by user"},
		// {Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (q *QueryCLI) execCommand(t string) {
	// insert the command into the history
	go q.H.Insert(t)

	// lex and parse the input
	l := lexer.New([]rune("?- " + t))
	if bsrSet, errs := parser.Parse(l); len(errs) > 0 {
		log.Println(errs)
	} else {
		a := ast.BuildStatementList(bsrSet.GetRoot())
		output := make(chan *resolver.Bindings, 1)
		log.Println("Resolving...")
		go q.R.ResolveStatementList(a, &resolver.Bindings{}, output)
		for v := range output {
			if v.Empty() {
				fmt.Println("Yes.")
			} else {
				fmt.Println("OUTPUT", v)
			}

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
		q.execCommand,
		completer,
		prompt.OptionPrefix("?- "),
		prompt.OptionSetExitCheckerOnInput(exitChecker),
		prompt.OptionHistory(q.H.Items()),
	)
	// Enter a REPL
	qPrompt.Run()

	return nil
}
