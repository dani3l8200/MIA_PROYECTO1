%{
// Copyright (c) 2015 The golex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is an example program using golex run time library.
package LWH

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

digit         ("-")?[0-9]+
identifier    ([a-zA-Z_])[a-zA-Z0-9_"-""."]*
letter        [_a-zA-Z]|{unicodeLetter}
id			 ([a-zA-Z_])[a-zA-Z0-9_"-""."]*
route 		 ("/"{id})("/"{id})*("/"|"."{id})?
stringL      ("\""[^"\""]*"\"")
newLine    	[\\][\*]
unicodeDigit  \x81 //ESTO ES PORQUE GO TRABAJA CON UTF-8 POR ESO ES NECESARIO 
unicodeLetter \x80
%%
	c = l.Rule0()

[ \t\r\n]+
pause lval.token = string(l.TokenBytes(nil)); return PAUSE
exec lval.token = string(l.TokenBytes(nil)); return EXEC
-path lval.token = string(l.TokenBytes(nil)); return PATH 
-> lval.token = string(l.TokenBytes(nil)); return ARROW 
-unit  lval.token = string(l.TokenBytes(nil)); return UNIT
mkdisk lval.token = string(l.TokenBytes(nil)); return MKDISK
-size lval.token = string(l.TokenBytes(nil)); return SIZE
-name lval.token = string(l.TokenBytes(nil)); return NAME 
k lval.token = string(l.TokenBytes(nil)); return K
m lval.token = string(l.TokenBytes(nil)); return M 
rmdisk    lval.token = string(l.TokenBytes(nil)); return RMDISK
-type lval.token = string(l.TokenBytes(nil)); return TYPE
b lval.token = string(l.TokenBytes(nil)); return B
p lval.token = string(l.TokenBytes(nil)); return P 
e lval.token = string(l.TokenBytes(nil)); return E 
l lval.token = string(l.TokenBytes(nil)); return L 
-fit lval.token = string(l.TokenBytes(nil)); return FIT 
bf lval.token = string(l.TokenBytes(nil)); return BF
ff lval.token = string(l.TokenBytes(nil)); return FF 
wf lval.token = string(l.TokenBytes(nil)); return WF
-delete lval.token = string(l.TokenBytes(nil)); return DELETE 
fast lval.token = string(l.TokenBytes(nil)); return FAST
full lval.token = string(l.TokenBytes(nil)); return FULL
-add lval.token = string(l.TokenBytes(nil)); return ADD
fdisk lval.token = string(l.TokenBytes(nil)); return FDISK 
{digit}  lval.token = string(l.TokenBytes(nil)); return NUMBERN
{id}	lval.token = string(l.TokenBytes(nil)); return ID
{stringL} lval.token = string(l.TokenBytes(nil)); return STRTYPE
{route} lval.token = string(l.TokenBytes(nil)); return ROUTE
{newLine} lval.token = string(l.TokenBytes(nil)); return CONTINUE
%%
	if c, ok := l.Abort(); ok { return int(c) }
	goto yyAction
}