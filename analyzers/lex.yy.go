// Code generated by golex. DO NOT EDIT.

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
	if err != nil {
		panic(err)
	}
	return &lexer{lx}
}

func (l *lexer) Lex(lval *yySymType) int {
	c := l.Enter()

yystate0:
	yyrule := -1
	_ = yyrule
	c = l.Rule0()

	goto yystart1

yyAction:
	switch yyrule {
	case 1:
		goto yyrule1
	case 2:
		goto yyrule2
	case 3:
		goto yyrule3
	case 4:
		goto yyrule4
	case 5:
		goto yyrule5
	case 6:
		goto yyrule6
	case 7:
		goto yyrule7
	case 8:
		goto yyrule8
	case 9:
		goto yyrule9
	case 10:
		goto yyrule10
	case 11:
		goto yyrule11
	case 12:
		goto yyrule12
	case 13:
		goto yyrule13
	case 14:
		goto yyrule14
	case 15:
		goto yyrule15
	case 16:
		goto yyrule16
	case 17:
		goto yyrule17
	case 18:
		goto yyrule18
	case 19:
		goto yyrule19
	case 20:
		goto yyrule20
	case 21:
		goto yyrule21
	case 22:
		goto yyrule22
	case 23:
		goto yyrule23
	case 24:
		goto yyrule24
	case 25:
		goto yyrule25
	case 26:
		goto yyrule26
	case 27:
		goto yyrule27
	case 28:
		goto yyrule28
	case 29:
		goto yyrule29
	case 30:
		goto yyrule30
	case 31:
		goto yyrule31
	case 32:
		goto yyrule32
	case 33:
		goto yyrule33
	case 34:
		goto yyrule34
	case 35:
		goto yyrule35
	case 36:
		goto yyrule36
	case 37:
		goto yyrule37
	case 38:
		goto yyrule38
	case 39:
		goto yyrule39
	case 40:
		goto yyrule40
	case 41:
		goto yyrule41
	case 42:
		goto yyrule42
	case 43:
		goto yyrule43
	case 44:
		goto yyrule44
	case 45:
		goto yyrule45
	case 46:
		goto yyrule46
	case 47:
		goto yyrule47
	}
yystate1:
	c = l.Next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate3
	case c == '#':
		goto yystate5
	case c == '-':
		goto yystate7
	case c == '/':
		goto yystate57
	case c == '=':
		goto yystate60
	case c == 'A' || c == 'C' || c == 'D' || c >= 'G' && c <= 'J' || c == 'N' || c == 'O' || c == 'Q' || c == 'S' || c == 'T' || c == 'V' || c >= 'X' && c <= 'Z' || c == '_' || c == 'a' || c == 'c' || c == 'd' || c >= 'g' && c <= 'j' || c == 'n' || c == 'o' || c == 'q' || c == 's' || c == 't' || c == 'v' || c >= 'x' && c <= 'z':
		goto yystate61
	case c == 'B' || c == 'b':
		goto yystate62
	case c == 'E' || c == 'e':
		goto yystate64
	case c == 'F' || c == 'f':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate80
	case c == 'L' || c == 'l':
		goto yystate81
	case c == 'M' || c == 'm':
		goto yystate89
	case c == 'P' || c == 'p':
		goto yystate105
	case c == 'R' || c == 'r':
		goto yystate110
	case c == 'U' || c == 'u':
		goto yystate118
	case c == 'W' || c == 'w':
		goto yystate125
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	case c >= '0' && c <= '9':
		goto yystate8
	}

yystate2:
	c = l.Next()
	yyrule = 1
	l.Mark()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate4
	case c >= '\x01' && c <= '!' || c >= '#' && c <= 'ÿ':
		goto yystate3
	}

yystate4:
	c = l.Next()
	yyrule = 44
	l.Mark()
	goto yyrule44

yystate5:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate6
	}

yystate6:
	c = l.Next()
	yyrule = 47
	l.Mark()
	switch {
	default:
		goto yyrule47
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate6
	}

yystate7:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '>':
		goto yystate9
	case c == 'A' || c == 'a':
		goto yystate10
	case c == 'C' || c == 'c':
		goto yystate13
	case c == 'D' || c == 'd':
		goto yystate17
	case c == 'F' || c == 'f':
		goto yystate23
	case c == 'I' || c == 'i':
		goto yystate26
	case c == 'N' || c == 'n':
		goto yystate29
	case c == 'P' || c == 'p':
		goto yystate33
	case c == 'R' || c == 'r':
		goto yystate39
	case c == 'S' || c == 's':
		goto yystate43
	case c == 'T' || c == 't':
		goto yystate47
	case c == 'U' || c == 'u':
		goto yystate51
	case c >= '0' && c <= '9':
		goto yystate8
	}

yystate8:
	c = l.Next()
	yyrule = 42
	l.Mark()
	switch {
	default:
		goto yyrule42
	case c >= '0' && c <= '9':
		goto yystate8
	}

yystate9:
	c = l.Next()
	yyrule = 5
	l.Mark()
	goto yyrule5

yystate10:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'D' || c == 'd':
		goto yystate11
	}

yystate11:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'D' || c == 'd':
		goto yystate12
	}

yystate12:
	c = l.Next()
	yyrule = 25
	l.Mark()
	goto yyrule25

yystate13:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'O' || c == 'o':
		goto yystate14
	}

yystate14:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'N' || c == 'n':
		goto yystate15
	}

yystate15:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'T' || c == 't':
		goto yystate16
	}

yystate16:
	c = l.Next()
	yyrule = 33
	l.Mark()
	goto yyrule33

yystate17:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'E' || c == 'e':
		goto yystate18
	}

yystate18:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'L' || c == 'l':
		goto yystate19
	}

yystate19:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'E' || c == 'e':
		goto yystate20
	}

yystate20:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'T' || c == 't':
		goto yystate21
	}

yystate21:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'E' || c == 'e':
		goto yystate22
	}

yystate22:
	c = l.Next()
	yyrule = 22
	l.Mark()
	goto yyrule22

yystate23:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'I' || c == 'i':
		goto yystate24
	}

yystate24:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'T' || c == 't':
		goto yystate25
	}

yystate25:
	c = l.Next()
	yyrule = 18
	l.Mark()
	goto yyrule18

yystate26:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'D' || c == 'd':
		goto yystate27
	}

yystate27:
	c = l.Next()
	yyrule = 30
	l.Mark()
	switch {
	default:
		goto yyrule30
	case c >= '0' && c <= '9':
		goto yystate28
	}

yystate28:
	c = l.Next()
	yyrule = 46
	l.Mark()
	switch {
	default:
		goto yyrule46
	case c >= '0' && c <= '9':
		goto yystate28
	}

yystate29:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'A' || c == 'a':
		goto yystate30
	}

yystate30:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'M' || c == 'm':
		goto yystate31
	}

yystate31:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'E' || c == 'e':
		goto yystate32
	}

yystate32:
	c = l.Next()
	yyrule = 9
	l.Mark()
	goto yyrule9

yystate33:
	c = l.Next()
	yyrule = 32
	l.Mark()
	switch {
	default:
		goto yyrule32
	case c == 'A' || c == 'a':
		goto yystate34
	case c == 'W' || c == 'w':
		goto yystate37
	}

yystate34:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'T' || c == 't':
		goto yystate35
	}

yystate35:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'H' || c == 'h':
		goto yystate36
	}

yystate36:
	c = l.Next()
	yyrule = 4
	l.Mark()
	goto yyrule4

yystate37:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'D' || c == 'd':
		goto yystate38
	}

yystate38:
	c = l.Next()
	yyrule = 38
	l.Mark()
	goto yyrule38

yystate39:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'U' || c == 'u':
		goto yystate40
	}

yystate40:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'T' || c == 't':
		goto yystate41
	}

yystate41:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'A' || c == 'a':
		goto yystate42
	}

yystate42:
	c = l.Next()
	yyrule = 41
	l.Mark()
	goto yyrule41

yystate43:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'I' || c == 'i':
		goto yystate44
	}

yystate44:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'Z' || c == 'z':
		goto yystate45
	}

yystate45:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'E' || c == 'e':
		goto yystate46
	}

yystate46:
	c = l.Next()
	yyrule = 8
	l.Mark()
	goto yyrule8

yystate47:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'Y' || c == 'y':
		goto yystate48
	}

yystate48:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'P' || c == 'p':
		goto yystate49
	}

yystate49:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'E' || c == 'e':
		goto yystate50
	}

yystate50:
	c = l.Next()
	yyrule = 13
	l.Mark()
	goto yyrule13

yystate51:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'N' || c == 'n':
		goto yystate52
	case c == 'S' || c == 's':
		goto yystate55
	}

yystate52:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'I' || c == 'i':
		goto yystate53
	}

yystate53:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'T' || c == 't':
		goto yystate54
	}

yystate54:
	c = l.Next()
	yyrule = 6
	l.Mark()
	goto yyrule6

yystate55:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == 'R' || c == 'r':
		goto yystate56
	}

yystate56:
	c = l.Next()
	yyrule = 37
	l.Mark()
	goto yyrule37

yystate57:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate58
	}

yystate58:
	c = l.Next()
	yyrule = 45
	l.Mark()
	switch {
	default:
		goto yyrule45
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate58
	case c == '/':
		goto yystate59
	}

yystate59:
	c = l.Next()
	yyrule = 45
	l.Mark()
	switch {
	default:
		goto yyrule45
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate58
	}

yystate60:
	c = l.Next()
	yyrule = 31
	l.Mark()
	goto yyrule31

yystate61:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate62:
	c = l.Next()
	yyrule = 14
	l.Mark()
	switch {
	default:
		goto yyrule14
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate61
	case c == 'F' || c == 'f':
		goto yystate63
	}

yystate63:
	c = l.Next()
	yyrule = 19
	l.Mark()
	switch {
	default:
		goto yyrule19
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate64:
	c = l.Next()
	yyrule = 16
	l.Mark()
	switch {
	default:
		goto yyrule16
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate61
	case c == 'X' || c == 'x':
		goto yystate65
	}

yystate65:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate61
	case c == 'E' || c == 'e':
		goto yystate66
	}

yystate66:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate61
	case c == 'C' || c == 'c':
		goto yystate67
	}

yystate67:
	c = l.Next()
	yyrule = 3
	l.Mark()
	switch {
	default:
		goto yyrule3
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate68:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c == 'B' || c == 'C' || c == 'E' || c >= 'G' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c == 'b' || c == 'c' || c == 'e' || c >= 'g' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate61
	case c == 'A' || c == 'a':
		goto yystate69
	case c == 'D' || c == 'd':
		goto yystate72
	case c == 'F' || c == 'f':
		goto yystate76
	case c == 'U' || c == 'u':
		goto yystate77
	}

yystate69:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate61
	case c == 'S' || c == 's':
		goto yystate70
	}

yystate70:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate61
	case c == 'T' || c == 't':
		goto yystate71
	}

yystate71:
	c = l.Next()
	yyrule = 23
	l.Mark()
	switch {
	default:
		goto yyrule23
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate72:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate61
	case c == 'I' || c == 'i':
		goto yystate73
	}

yystate73:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate61
	case c == 'S' || c == 's':
		goto yystate74
	}

yystate74:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate61
	case c == 'K' || c == 'k':
		goto yystate75
	}

yystate75:
	c = l.Next()
	yyrule = 26
	l.Mark()
	switch {
	default:
		goto yyrule26
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate76:
	c = l.Next()
	yyrule = 20
	l.Mark()
	switch {
	default:
		goto yyrule20
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate77:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate61
	case c == 'L' || c == 'l':
		goto yystate78
	}

yystate78:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate61
	case c == 'L' || c == 'l':
		goto yystate79
	}

yystate79:
	c = l.Next()
	yyrule = 24
	l.Mark()
	switch {
	default:
		goto yyrule24
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate80:
	c = l.Next()
	yyrule = 10
	l.Mark()
	switch {
	default:
		goto yyrule10
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate81:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate61
	case c == 'O' || c == 'o':
		goto yystate82
	}

yystate82:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate61
	case c == 'G' || c == 'g':
		goto yystate83
	}

yystate83:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate61
	case c == 'I' || c == 'i':
		goto yystate84
	case c == 'O' || c == 'o':
		goto yystate86
	}

yystate84:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate61
	case c == 'N' || c == 'n':
		goto yystate85
	}

yystate85:
	c = l.Next()
	yyrule = 36
	l.Mark()
	switch {
	default:
		goto yyrule36
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate86:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate61
	case c == 'U' || c == 'u':
		goto yystate87
	}

yystate87:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate61
	case c == 'T' || c == 't':
		goto yystate88
	}

yystate88:
	c = l.Next()
	yyrule = 39
	l.Mark()
	switch {
	default:
		goto yyrule39
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate89:
	c = l.Next()
	yyrule = 11
	l.Mark()
	switch {
	default:
		goto yyrule11
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate61
	case c == 'K' || c == 'k':
		goto yystate90
	case c == 'O' || c == 'o':
		goto yystate101
	}

yystate90:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c == 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c == 'e' || c >= 'g' && c <= 'z':
		goto yystate61
	case c == 'D' || c == 'd':
		goto yystate91
	case c == 'F' || c == 'f':
		goto yystate96
	}

yystate91:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate61
	case c == 'I' || c == 'i':
		goto yystate92
	}

yystate92:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 't' && c <= 'z':
		goto yystate61
	case c == 'R' || c == 'r':
		goto yystate93
	case c == 'S' || c == 's':
		goto yystate94
	}

yystate93:
	c = l.Next()
	yyrule = 35
	l.Mark()
	switch {
	default:
		goto yyrule35
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate94:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate61
	case c == 'K' || c == 'k':
		goto yystate95
	}

yystate95:
	c = l.Next()
	yyrule = 7
	l.Mark()
	switch {
	default:
		goto yyrule7
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate96:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate61
	case c == 'I' || c == 'i':
		goto yystate97
	case c == 'S' || c == 's':
		goto yystate100
	}

yystate97:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate61
	case c == 'L' || c == 'l':
		goto yystate98
	}

yystate98:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate61
	case c == 'E' || c == 'e':
		goto yystate99
	}

yystate99:
	c = l.Next()
	yyrule = 34
	l.Mark()
	switch {
	default:
		goto yyrule34
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate100:
	c = l.Next()
	yyrule = 29
	l.Mark()
	switch {
	default:
		goto yyrule29
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate101:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate61
	case c == 'U' || c == 'u':
		goto yystate102
	}

yystate102:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate61
	case c == 'N' || c == 'n':
		goto yystate103
	}

yystate103:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate61
	case c == 'T' || c == 't':
		goto yystate104
	}

yystate104:
	c = l.Next()
	yyrule = 27
	l.Mark()
	switch {
	default:
		goto yyrule27
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate105:
	c = l.Next()
	yyrule = 15
	l.Mark()
	switch {
	default:
		goto yyrule15
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate61
	case c == 'A' || c == 'a':
		goto yystate106
	}

yystate106:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate61
	case c == 'U' || c == 'u':
		goto yystate107
	}

yystate107:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate61
	case c == 'S' || c == 's':
		goto yystate108
	}

yystate108:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate61
	case c == 'E' || c == 'e':
		goto yystate109
	}

yystate109:
	c = l.Next()
	yyrule = 2
	l.Mark()
	switch {
	default:
		goto yyrule2
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate110:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate61
	case c == 'E' || c == 'e':
		goto yystate111
	case c == 'M' || c == 'm':
		goto yystate113
	}

yystate111:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate61
	case c == 'P' || c == 'p':
		goto yystate112
	}

yystate112:
	c = l.Next()
	yyrule = 40
	l.Mark()
	switch {
	default:
		goto yyrule40
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate113:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate61
	case c == 'D' || c == 'd':
		goto yystate114
	}

yystate114:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate61
	case c == 'I' || c == 'i':
		goto yystate115
	}

yystate115:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate61
	case c == 'S' || c == 's':
		goto yystate116
	}

yystate116:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate61
	case c == 'K' || c == 'k':
		goto yystate117
	}

yystate117:
	c = l.Next()
	yyrule = 12
	l.Mark()
	switch {
	default:
		goto yyrule12
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate118:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate61
	case c == 'N' || c == 'n':
		goto yystate119
	}

yystate119:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate61
	case c == 'M' || c == 'm':
		goto yystate120
	}

yystate120:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate61
	case c == 'O' || c == 'o':
		goto yystate121
	}

yystate121:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate61
	case c == 'U' || c == 'u':
		goto yystate122
	}

yystate122:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate61
	case c == 'N' || c == 'n':
		goto yystate123
	}

yystate123:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate61
	case c == 'T' || c == 't':
		goto yystate124
	}

yystate124:
	c = l.Next()
	yyrule = 28
	l.Mark()
	switch {
	default:
		goto yyrule28
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yystate125:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate61
	case c == 'F' || c == 'f':
		goto yystate126
	}

yystate126:
	c = l.Next()
	yyrule = 21
	l.Mark()
	switch {
	default:
		goto yyrule21
	case c == '"' || c == '.' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate61
	}

yyrule1: // [ \t\r\n]+

	goto yystate0
yyrule2: // pause
	{
		lval.token = string(l.TokenBytes(nil))
		return PAUSE
		goto yystate0
	}
yyrule3: // exec
	{
		lval.token = string(l.TokenBytes(nil))
		return EXEC
		goto yystate0
	}
yyrule4: // -path
	{
		lval.token = string(l.TokenBytes(nil))
		return PATH
		goto yystate0
	}
yyrule5: // ->
	{
		lval.token = string(l.TokenBytes(nil))
		return ARROW
		goto yystate0
	}
yyrule6: // -unit
	{
		lval.token = string(l.TokenBytes(nil))
		return UNIT
		goto yystate0
	}
yyrule7: // mkdisk
	{
		lval.token = string(l.TokenBytes(nil))
		return MKDISK
		goto yystate0
	}
yyrule8: // -size
	{
		lval.token = string(l.TokenBytes(nil))
		return SIZE
		goto yystate0
	}
yyrule9: // -name
	{
		lval.token = string(l.TokenBytes(nil))
		return NAME
		goto yystate0
	}
yyrule10: // k
	{
		lval.token = string(l.TokenBytes(nil))
		return K
		goto yystate0
	}
yyrule11: // m
	{
		lval.token = string(l.TokenBytes(nil))
		return M
		goto yystate0
	}
yyrule12: // rmdisk
	{
		lval.token = string(l.TokenBytes(nil))
		return RMDISK
		goto yystate0
	}
yyrule13: // -type
	{
		lval.token = string(l.TokenBytes(nil))
		return TYPE
		goto yystate0
	}
yyrule14: // b
	{
		lval.token = string(l.TokenBytes(nil))
		return B
		goto yystate0
	}
yyrule15: // p
	{
		lval.token = string(l.TokenBytes(nil))
		return P
		goto yystate0
	}
yyrule16: // e
	{
		lval.token = string(l.TokenBytes(nil))
		return E
		goto yystate0
	}
yyrule17: // l
	{
		lval.token = string(l.TokenBytes(nil))
		return L
		goto yystate0
	}
yyrule18: // -fit
	{
		lval.token = string(l.TokenBytes(nil))
		return FIT
		goto yystate0
	}
yyrule19: // bf
	{
		lval.token = string(l.TokenBytes(nil))
		return BF
		goto yystate0
	}
yyrule20: // ff
	{
		lval.token = string(l.TokenBytes(nil))
		return FF
		goto yystate0
	}
yyrule21: // wf
	{
		lval.token = string(l.TokenBytes(nil))
		return WF
		goto yystate0
	}
yyrule22: // -delete
	{
		lval.token = string(l.TokenBytes(nil))
		return DELETE
		goto yystate0
	}
yyrule23: // fast
	{
		lval.token = string(l.TokenBytes(nil))
		return FAST
		goto yystate0
	}
yyrule24: // full
	{
		lval.token = string(l.TokenBytes(nil))
		return FULL
		goto yystate0
	}
yyrule25: // -add
	{
		lval.token = string(l.TokenBytes(nil))
		return ADD
		goto yystate0
	}
yyrule26: // fdisk
	{
		lval.token = string(l.TokenBytes(nil))
		return FDISK
		goto yystate0
	}
yyrule27: // mount
	{
		lval.token = string(l.TokenBytes(nil))
		return MOUNT
		goto yystate0
	}
yyrule28: // unmount
	{
		lval.token = string(l.TokenBytes(nil))
		return UNMOUNT
		goto yystate0
	}
yyrule29: // mkfs
	{
		lval.token = string(l.TokenBytes(nil))
		return MKFS
		goto yystate0
	}
yyrule30: // -id
	{
		lval.token = string(l.TokenBytes(nil))
		return IDN
		goto yystate0
	}
yyrule31: // =
	{
		lval.token = string(l.TokenBytes(nil))
		return S_EQUAL
		goto yystate0
	}
yyrule32: // -p
	{
		lval.token = string(l.TokenBytes(nil))
		return PCONT
		goto yystate0
	}
yyrule33: // -cont
	{
		lval.token = string(l.TokenBytes(nil))
		return CONT
		goto yystate0
	}
yyrule34: // mkfile
	{
		lval.token = string(l.TokenBytes(nil))
		return MKFILE
		goto yystate0
	}
yyrule35: // mkdir
	{
		lval.token = string(l.TokenBytes(nil))
		return MKDIR
		goto yystate0
	}
yyrule36: // login
	{
		lval.token = string(l.TokenBytes(nil))
		return LOGIN
		goto yystate0
	}
yyrule37: // -usr
	{
		lval.token = string(l.TokenBytes(nil))
		return USER
		goto yystate0
	}
yyrule38: // -pwd
	{
		lval.token = string(l.TokenBytes(nil))
		return PWD
		goto yystate0
	}
yyrule39: // logout
	{
		lval.token = string(l.TokenBytes(nil))
		return LOGOUT
		goto yystate0
	}
yyrule40: // rep
	{
		lval.token = string(l.TokenBytes(nil))
		return REP
		goto yystate0
	}
yyrule41: // -ruta
	{
		lval.token = string(l.TokenBytes(nil))
		return RUTA
		goto yystate0
	}
yyrule42: // {digit}
	{
		lval.token = string(l.TokenBytes(nil))
		return NUMBERN
		goto yystate0
	}
yyrule43: // {id}
	{
		lval.token = string(l.TokenBytes(nil))
		return ID
		goto yystate0
	}
yyrule44: // {stringL}
	{
		lval.token = string(l.TokenBytes(nil))
		return STRTYPE
		goto yystate0
	}
yyrule45: // {route}
	{
		lval.token = string(l.TokenBytes(nil))
		return ROUTE
		goto yystate0
	}
yyrule46: // {idn}
	{
		lval.token = string(l.TokenBytes(nil))
		return IDM
		goto yystate0
	}
yyrule47: // {comment}
	if true { // avoid go vet determining the below panic will not be reached
		lval.token = string(l.TokenBytes(nil))
		return COMMENT
		goto yystate0
	}
	panic("unreachable")

yyabort: // no lexem recognized
	//
	// silence unused label errors for build and satisfy go vet reachability analysis
	//
	{
		if false {
			goto yyabort
		}
		if false {
			goto yystate0
		}
		if false {
			goto yystate1
		}
	}

	if c, ok := l.Abort(); ok {
		return int(c)
	}
	goto yyAction
}
