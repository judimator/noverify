package main

import (
	"testing"

	"github.com/VKCOM/noverify/src/linttest"
)

func TestAssignmentAsExpression(t *testing.T) {
	test := linttest.NewSuite(t)

	test.AddFile(`<?php
	// phpdoc annotations are not required for NoVerify in simple cases
	function something() {
		$a = 'test';
		return $a;
	}
	function in_array() {}

	function test() {
		$b = ['1', '2', '3'];

		if (in_array(something(), $b)) {
			echo 'third arg true';
		}

		if (something() == $b[1]) {
			echo 'must be ===';
		}
	}`)

	test.Expect = []string{
		"3rd argument of in_array must be true when comparing strings",
		"Strings must be compared using '===' operator",
	}

	test.RunAndMatch()
}

func TestSingeQuotedString(t *testing.T) {
	test := linttest.NewSuite(t)

	test.AddFile(`<?php
$var0 = "string".'other string';
$var1 = "string";
$var2 = "other $var1";
$var3 = 'another';
$var4 = "another {$var3}";
`)

	test.Expect = []string{
		"Strings must be declared using ''",
		"Strings must be declared using ''",
	}

	test.RunAndMatch()
}
