package main

import (
	"log"
	"os"
)

const InFile string = "main.go"

func main() {
	log.Println("Booting Parsertongue")
	var (
		g                 = VerifyGrammar(OpenGrammar())
		file, _           = os.Open(InFile)
		lexToParseChan    = make(chan Token)
		lexer          *L = NewLexer(g, file, lexToParseChan)
	)
	log.Println("Grammar found.")
	log.Println("Starting to run the lexer")
	lexer.run()

	// Generate the pre-built Expressions first.
	// For each expression in the grammar that's nil,
	// make sure that it has the name of a pre-built expression.

	// Then, add it to the PrebuiltLexemes list.
	// The pre-built lexemes list is a map from the string to it's stateFn
	// This list is a global, and so if your production p.Expr is ever nil, then you need to look up it's name in the pre-built list and call it's stateFn.

	// Define a lexemes list global. It containers each of the lexeme productions:
	// 		lexeme production stuct {
	//      	ebnf.Production
	//          if it's nil, find the
	// 		}

	log.Println("Grammar successfully verified")
}
