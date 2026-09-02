# regex-parser

<!-- README: This project is fueled by curiosity and snacks. -->
An implementation of a regex parser in golang ^_^

This project currently supports tokenizing and parsing basic regular expressions into an Abstract Syntax Tree (AST).

## Current Features:

- Literal Characters
- Wildcard: .
- Quantifiers: *, +, ?
- Grouping with ()
- Alternation with |
- Implicit concatenation
- Escaped regex metacharacters
- AST pretty printing
- CLI for viewing tokens and parsed ASTs

## Example:

Parse a regex:

``` go run ./cmd/regex parse 'a(b|c)*' ```

Output:

```
Concat
  Literal(a)
  Star
    Alternation
      Literal(b)
      Literal(c)
```

View lexer tokens:

``` go run ./cmd/regex tokens 'a(b|c)*' ```

Output:

```
LITERAL(a) 
LPAREN 
LITERAL(b) 
PIPE 
LITERAL(c) 
RPAREN 
STAR 
EOF
```

Architecture:

Regex -> Lexer -> Tokens -> Parser -> AST -> NFA -> Matcher

The lexer, parser and AST are implemented.
The AST -> NFA compiler is implemented and tested; matching is available via the `matcher` package.

Made with <3 by Three Thirds!