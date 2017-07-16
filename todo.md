# Test the Lex Package

- Generate the coverage profile for the lexer
- Find the uncovered functions
- Add tests (including fixtures)
- Add the top level code to lex the entrant prods from the DAG

- Refactor the tests to DRY up reused patterns
- Switch to using -1 for failed matches instead of 0
    This will remove a bug in Sequences where optionals are
    nested inside of groups.

    ( [ nonterminal ] )

    This, for example, will be marked as non-matching
    when nonterminal is not matched, but it should match.

# Migrate the Lexer into the lex package

- Test the `FileLocation.String()` function
- Add helper functions to the token type.
- Add helper functions to the match type
- Test very carefully
- Build the DAG
    - The DAG consists of nodes:
        - Each nodes contains the production itself
        - And the array of state functions needed to match the body
- Clarify that TokenType type is a real type and stringifies correctly...
- Write an "AdvanceWhitespace" helper function.
Write up the low level helper functions for the lexer

# Correct the imports in the parsing package

# Update the toolchain to use Go 1.9 beta 

This is to allow the use of sync.Map and type aliasing.

# Implement the main.go function
