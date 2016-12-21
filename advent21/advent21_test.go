package main

import (
	"testing"
)

func TestReverseRunes(t *testing.T) {
	input := "abcdefg"
	output := string(reverseRunes([]rune(input)))
	if output != "gfedcba" {
		t.Error("Reverse of ", input, " returned ", output, ".")
	}
}

func TestRotateRightOnce(t *testing.T) {
	input := "abcdefg"
	output := rotateRight(input, 1)
	if output != "gabcdef" {
		t.Error("Single rotation right of ", input, " returned ", output, ".")
	}
}

func TestRotateRightThrice(t *testing.T) {
	input := "abcdefg"
	output := rotateRight(input, 3)
	if output != "efgabcd" {
		t.Error("Triple rotation right of ", input, " returned ", output, ".")
	}
}

func TestRotateLeftOnce(t *testing.T) {
	input := "abcdefg"
	output := rotateLeft(input, 1)
	if output != "bcdefga" {
		t.Error("Single rotation left of ", input, " returned ", output, ".")
	}
}

func TestRotateLeftThrice(t *testing.T) {
	input := "abcdefg"
	output := rotateLeft(input, 3)
	if output != "defgabc" {
		t.Error("Triple rotation left of ", input, " returned ", output, ".")
	}
}

func TestMoveLetterForward(t *testing.T) {
	input := "abcdefg"
	output := moveLetter(input, 2, 5)
	if output != "abdefcg" {
		t.Error("Move letter 2 to 5 of ", input, " returned ", output, ".")
	}
}

func TestMoveLetterBackward(t *testing.T) {
	input := "abcdefg"
	output := moveLetter(input, 5, 2)
	if output != "abfcdeg" {
		t.Error("Move letter 5 to 2 of ", input, " returned ", output, ".")
	}
}
