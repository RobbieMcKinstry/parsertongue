# Add to the Grammar package:

- Add the Verify Grammar function to the package
- Test the Verify Grammar package
- Test the grammar.init function
- Ensure that the grammar.init function effectively
    calls Verify Grammar

# Migrate the Lexer into the lex package

- Test the `FileLocation.String()` function
- Add helper functions to the token type.
- Add helper functions to the match type
- Test very carefully
- Clarify that TokenType type is a real type and stringifies correctly...
Write up the low level helper functions for the lexer


Add them together to get progressively closer to completing
the lexer

Test everything as you go

# Correct the imports in the parsing package

# Update the toolchain to use Go 1.9 beta 

This is to allow the use of sync.Map and type aliasing.

# Implement the main.go function
