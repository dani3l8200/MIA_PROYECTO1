%{
// Copyright (c) 2015 The golex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is an example program using golex run time library.
package analyzers

import (
  "bufio"
  "go/token"
  "io"
	"unicode"

	"modernc.org/golex/lex"
)

// Allocate Character classes anywhere in [0x80, 0xFF].
const (
	classUnicodeLeter = iota + 0x80
	classUnicodeDigit
	classOther
)

type lexer struct {
	*lex.Lexer
}

func rune2Class(r rune) int {
	if r >= 0 && r < 0x80 { // Keep ASCII as it is.
		return int(r)
	}
	if unicode.IsLetter(r) {
		return classUnicodeLeter
	}
	if unicode.IsDigit(r) {
		return classUnicodeDigit
	}
	return classOther
}

func newLexer(src io.Reader, dst io.Writer, fName string) *lexer {
  file := token.NewFileSet().AddFile(fName, -1, 1<<31-1)
  lx, err := lex.New(file, bufio.NewReader(src), lex.RuneClass(rune2Class))
  if (err != nil) { panic(err) }
  return &lexer{lx}
}

func (l *lexer) Lex(lval *yySymType) int {
  c := l.Enter()

%}
%option case-insensitive
%yyc c
%yyn c = l.Next()
%yym l.Mark()

digit         [0-9]|{unicodeDigit}
identifier    {letter}({letter}|{digit})*
letter        [_a-zA-Z]|{unicodeLetter}
unicodeDigit  \x81
unicodeLetter \x80
op            [-+*/]
%%
	c = l.Rule0()

[ \t\r\n]+
int  lval.token = string(l.TokenBytes(nil)); return INT
string lval.token = string(l.TokenBytes(nil)); return STRING
float lval.token = string(l.TokenBytes(nil)); return FLOAT
bool lval.token = string(l.TokenBytes(nil)); return BOOL
char lval.token = string(l.TokenBytes(nil)); return CHAR
{identifier} lval.token = string(l.TokenBytes(nil)); return IDENT
%%
	if c, ok := l.Abort(); ok { return int(c) }
	goto yyAction
}