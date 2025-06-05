package main

import (
	"go_reloaded/text_processing"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    string
	}{
		// Basic Conversions
		{
			"Hex conversion",
			"1E (hex) files were added",
			"30 files were added",
		},
		{
			"Hex conversion with lowercase",
			"a(hex)",
			"10",
		},
		{
			"Binary conversion",
			"It has been 10 (bin) years",
			"It has been 2 years",
		},
		{
			"Uppercase transformation",
			"Ready, set, go (up, 3) !",
			"READY, SET, GO!",
		},
		{
			"Uppercase transformation and punctuation with spaces",
			"Ready, set, go (up) !",
			"Ready, set, GO!",
		},
		{
			"Lowercase transformation",
			"I should stop SHOUTING (low)",
			"I should stop shouting",
		},
		{
			"Capitalize transformation",
			"Welcome to the Brooklyn bridge (cap)",
			"Welcome to the Brooklyn Bridge",
		},
		{
			"Uppercase specific words",
			"This is so exciting (up, 2)",
			"This is SO EXCITING",
		},
		{
			"Multiple conversions combined",
			"Simply add 42 (hex) and 10 (bin)",
			"Simply add 66 and 2",
		},
		// Article Corrections
		{
			"Article correction before vowel in quotes",
			"A 'apple'",
			"An 'apple'",
		},
		{
			"Article correction before vowel in quotes",
			"A 'heiress'",
			"An 'heiress'",
		},
		{
			"Article correction before vowel",
			"A apple",
			"An apple",
		},
		{
			"Article correction before vowel (lowercase)",
			"a apple",
			"an apple",
		},
		{
			"Article correction before consonant",
			"an cat an hat an dog",
			"a cat a hat a dog",
		},
		{
			"Article correction before 'h' sounds",
			"a hour a heir a honor a honest man",
			"an hour an heir an honor an honest man",
		},
		{
			"Article correction with multiple instances",
			" ' There is no greater agony than bearing a untold story . '",
			"'There is no greater agony than bearing an untold story.'",
		},
		{
			"Ignore article correction when only a",
			"a a a ",
			"a a a",
		},
		{
			"Ignore article correction with 'or'",
			"a or b",
			"a or b",
		},
		{
			"Ignore article correction with 'and'",
			"a and the",
			"a and the",
		},
		{
			"Article correction in quoted text",
			"I am a optimist, but a optimist",
			"I am an optimist, but an optimist",
		},
		// Multiword Modifications
		{
			"Order of modifications",
			"a (cap) alem (cap) and a alem (up, 2)",
			"An Alem and AN ALEM",
		},
		{
			"Order of modifications",
			"a (up) ALEM (low) a alem (up, 3)",
			"AN ALEM AN ALEM",
		},
		{
			"Capitalize last 2 words",
			"harold wilson (cap, 2)",
			"Harold Wilson",
		},
		{
			"Capitalize last 6 words",
			"it was the age of wisdom, 'it was the age of foolishness' (cap, 6)",
			"it was the age of wisdom, 'It Was The Age Of Foolishness'",
		},
		{
			"Lowercase last 3 words",
			"IT WAS THE (low, 3) winter of despair",
			"it was the winter of despair",
		},
		{
			"Uppercase last 2 words",
			"did you get in my house (up, 2)",
			"did you get in MY HOUSE",
		},
		{
			"Mixed case transformation sequence",
			"it (cap) was the best of times, it was the worst of times (up)",
			"It was the best of times, it was the worst of TIMES",
		},
		// Punctuation Handling
		{
			"Comma spacing",
			"hello,there",
			"hello, there",
		},
		{
			"Multiple spaces reduction",
			"I was sitting over    !? . there",
			"I was sitting over!?. there",
		},
		{
			"Ellipsis from multiple periods",
			"I was thinking .  .    ........... .....",
			"I was thinking..................",
		},
		{
			"Terminal punctuation spacing",
			"There it was. A amazing rock!",
			"There it was. An amazing rock!",
		},
		{
			"Multiple exclamation marks",
			"BAMM !  !  !!!!",
			"BAMM!!!!!!",
		},
		{
			"Mixed punctuation cleanup",
			"Punctuation tests are:::: ;;; ;; ????? ,,,,,,",
			"Punctuation tests are::::;;;;;?????,,,,,,",
		},
		{
			"Double dots",
			"Puncuation tests are .   .   ",
			"Puncuation tests are..",
		},
		{
			"Period spacing",
			"das . And",
			"das. And",
		},

		// Quote Handling
		{
			"Single quotes spacing",
			"'one''two''three'",
			"'one' 'two' 'three'",
		},
		{
			"Transforming outside quotes",
			"it (cap) was an 'great' (up) experience?!",
			"It was a 'GREAT' experience?!",
		},
		{
			"Single quotes spacing",
			"' awesome '",
			"'awesome'",
		},
		{
			"Quotes with transformation",
			"'Transform inside (up)'",
			"'Transform INSIDE'",
		},
		{
			"Quoted speech with attribution",
			"'she said. '",
			"'she said.'",
		},
		{
			"Quotes cleanup",
			"hi ' hi' hi",
			"hi 'hi' hi",
		},
		{
			"Quotes cleanup",
			" ' hi' hi'",
			"'hi' hi'",
		},
		{
			"Quotes cleanup",
			"hi 'hi",
			"hi 'hi",
		},

		// Edge Cases
		{
			"Big int conversion",
			"18284829229392927abcd17283930 (hex)",
			"7839508634210545570538414416607536",
		},
		{
			"Invalid hex remains unchanged",
			"(hex",
			"(hex",
		},
		{
			"Valide standalone transformation command",
			"(up)",
			"",
		},
		{
			"Malformed transformation",
			"(aaaa)",
			"(aaaa)",
		},
		{
			"String with multiple spaces",
			"Elton       John",
			"Elton John",
		},
		{
			"Leading spaces trimming",
			" wanna chose",
			"wanna chose",
		},
		{
			"Combined transformations",
			"one(low)two(cap)three(up)",
			"one Two THREE",
		},
		{
			"Preserving valid punctuation sequences",
			"BAMM!!!",
			"BAMM!!!",
		},
		{
			"Complex combined transformation",
			"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			"I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			"Transformation with zero count",
			"hello (up, 0) world",
			"hello world",
		},
		{
			"Invalid numeric transformation",
			"abg (hex)",
			"abg",
		},
		{
			"Transformation with negative count",
			"hello (up, -2) world",
			"hello world",
		},
		{
			"Sequential transformations on different words",
			"hello (up) world (cap)",
			"HELLO World",
		},
		{
			"Invalid numeric transformation transforming",
			"ab (Hex) (low)",
			"171",
		},
		{
			"Invalid numeric transformation transforming",
			"ab (up) (HeX) (low)",
			"171",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			output := text_processing.ProcessText(tt.input)

			if output != tt.expected {
				t.Errorf("%s:\n input: %q\n output %q\n wants: %q", tt.description, tt.input, output, tt.expected)
			}
		})
	}
}
