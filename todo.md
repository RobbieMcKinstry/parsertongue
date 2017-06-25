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

# Implement Parallel LL(1) Lexing

- Create a root containing the lexemes from the DAG.
- Test the DAG
- In parallel, get back the length of each lexeme recursively.
    - You don't need to worry about conflicts because the grammar is LL(1)
    - You only need to get back the width of each lexeme, because that will tell you how big the token is (and all you care about is reporting the top level token, not any of the intermediary tokens.

Add them together to get progressively closer to completing
the lexer

Test everything as you go

# Correct the imports in the parsing package

# Update the toolchain to use Go 1.9 beta 

This is to allow the use of sync.Map and type aliasing.

# Implement the main.go function
