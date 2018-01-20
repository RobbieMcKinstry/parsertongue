# Grammar Visualization

- Don't send the whole grammar. Send only the production names
and their children nodes.

- Add edges to the visualization so each production is linked to 
its child nodes.

- Encode the relationship around those nodes for visualizatoin

# Golang Progression

The goal is to make progress on the go lexing.
The next step is go deeper into lexing strings. Right now we lex half of
the string definition, and I need to go a little deeper.

- Add string fixtures for the full golang string definition

Next, grab the whole import definition

- Add a fixture for the ImportSpec definition

I need to get Clean working without panicking.

- Uncomment clean() calls and debug the nil pointer.

Then try to run the full golang spec

- Comment the Golang skip step.

# Test the Lex Package

- Generate the coverage profile for the lexer
- Find the uncovered functions
- Add tests (including fixtures)

- Refactor the tests to DRY up reused patterns

# Correct the imports in the parsing package

# Implement the main.go function
