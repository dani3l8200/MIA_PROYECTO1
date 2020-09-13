%{
// Copyright (c) 2015 The golex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is an example program using golex run time library.
// analyzers Contiene los analizadores usados para la lectura inicial
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
comment 	[#][^\n]+
digit         ("-")?[0-9]+
identifier    ([a-zA-Z_])[a-zA-Z0-9_"-""."]*
letter        [_a-zA-Z]|{unicodeLetter}
id			 ([a-zA-Z_])[a-zA-Z0-9_"-""."]*
route 		 ("/"{id})("/"{id})*("/"|"."{id})?
stringL      ("\""[^"\""]*"\"")
idn  		 (-id[0-9]+)	
unicodeDigit  \x81
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
mount lval.token = string(l.TokenBytes(nil)); return MOUNT
unmount lval.token = string(l.TokenBytes(nil)); return UNMOUNT
mkfs lval.token = string(l.TokenBytes(nil)); return MKFS 
-id lval.token = string(l.TokenBytes(nil)); return IDN
= lval.token = string(l.TokenBytes(nil)); return S_EQUAL
-p lval.token = string(l.TokenBytes(nil)); return PCONT
-cont lval.token = string(l.TokenBytes(nil)); return CONT
mkfile lval.token = string(l.TokenBytes(nil)); return MKFILE
mkdir lval.token = string(l.TokenBytes(nil)); return MKDIR
login lval.token = string(l.TokenBytes(nil)); return LOGIN
-usr lval.token = string(l.TokenBytes(nil)); return USER
-pwd lval.token = string(l.TokenBytes(nil)); return PWD
logout lval.token = string(l.TokenBytes(nil)); return LOGOUT
rep lval.token = string(l.TokenBytes(nil)); return REP
-ruta lval.token = string(l.TokenBytes(nil)); return RUTA
{digit}  lval.token = string(l.TokenBytes(nil)); return NUMBERN
{id}	lval.token = string(l.TokenBytes(nil)); return ID
{stringL} lval.token = string(l.TokenBytes(nil)); return STRTYPE
{route} lval.token = string(l.TokenBytes(nil)); return ROUTE
{idn} lval.token = string(l.TokenBytes(nil)); return IDM
{comment} lval.token = string(l.TokenBytes(nil)); return COMMENT
%%
	if c, ok := l.Abort(); ok { return int(c) }
	goto yyAction
}