package additional_functions

import "regexp"

var (
	Checking_Punctuation = regexp.MustCompile(`^[.,!?;:]+$`)
	IsHexCheck = "^[0-9a-fA-F]+$"
	RegToken = regexp.MustCompile(`\([a-zA-Z]+(?:,\s*-?\d+)?\)|(\([^)]*\))|(\([^)]*)|[\w']+|[.,!?;:]+|\n`)
	RemoveSpaceBeforePunct = regexp.MustCompile(`\s+([.,!?;:])`)
	RemoveSpaceBetweenPuncts = regexp.MustCompile(`([.,!?;:])\s+([.,!?;:])`)
	ApostropheContentWithSpaces = regexp.MustCompile(`'(\s*)([^']*?)(\s*)'`)
	SplitAdjacentApostrophes = regexp.MustCompile(`'([^']*)''([^']*)'`)

)
