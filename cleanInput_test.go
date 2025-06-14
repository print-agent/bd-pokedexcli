package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	type runTests []struct {
		input    string
		expected []string
	}

	tests := runTests{
		{
			" Hello, World!  ",
			[]string{"Hello", "World"},
		},
		{
			"You have within you, right now, everything you need to deal with the world",
			[]string{"You", "Have", "Within", "You", "Right", "Now", "Everything", "You", "Need", "To", "Deal", "With", "The", "World"},
		},
		{
			"Nothing is impossible, the word itself says 'I'm impossible'!",
			[]string{"Nothing", "Is", "Impossible", "The", "Word", "Itself", "Says", "I'm", "Impossible"},
		},
		{
			"The quick brown fox jumps over the lazy dog",
			[]string{"The", "Quick", "Brown", "Fox", "Jumps", "Over", "The", "Lazy", "Dog"},
		},
		{
			"  Multiple   spaces   between   words   should   be   handled  ",
			[]string{"Multiple", "Spaces", "Between", "Words", "Should", "Be", "Handled"},
		},
		{
			"Short text",
			[]string{"Short", "Text"},
		},
		{
			"One",
			[]string{"One"},
		},
		{
			"",
			[]string{},
		},
		{
			"   ",
			[]string{},
		},
		{
			"Success is not final, failure is not fatal: it is the courage to continue that counts",
			[]string{"Success", "Is", "Not", "Final", "Failure", "Is", "Not", "Fatal", "It", "Is", "The", "Courage", "To", "Continue", "That", "Counts"},
		},
		{
			"UPPERCASE and lowercase MiXeD cAsE text",
			[]string{"Uppercase", "And", "Lowercase", "Mixed", "Case", "Text"},
		},
		{
			"Numbers123 and symbols!@# mixed with words",
			[]string{"Numbers123", "And", "Symbols", "Mixed", "With", "Words"},
		},
		{
			"Punctuation, marks; should: be! removed? from words.",
			[]string{"Punctuation", "Marks", "Should", "Be", "Removed", "From", "Words"},
		},
		{
			"a",
			[]string{"A"},
		},
		{
			"This is a longer sentence with many words to test the complete extraction functionality",
			[]string{"This", "Is", "A", "Longer", "Sentence", "With", "Many", "Words", "To", "Test", "The", "Complete", "Extraction", "Functionality"},
		},
	}

	for _, c := range tests {
		actual := cleanInput(c.input)

		fmt.Println(len(actual))
		fmt.Println(len(c.expected))

		if len(actual) != len(c.expected) {
			t.Errorf("Fail, length dont match")
		}

		// for i, word := range actual {
		// 	expectedWord := tCase.expected
		//
		// 	if word != expectedWord[i] {
		// 		t.Errorf("Fail, words dont match")
		// 	}
		// }
	}
}
