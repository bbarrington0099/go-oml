package gooml

import (
	"fmt"
	"slices"
)

// genDecoratorCombinations generates decorator combinations with a max of maxCombo decorators
func genDecoratorCombinations(decorators []*oml, maxCombo int) []struct {
	keyword string
	literal string
	ref     string
} {
	result := make([]struct {
		keyword string
		literal string
		ref     string
	}, 0)

	n := len(decorators)
	// Generate combinations with 1 to maxCombo decorators
	for size := 1; size <= maxCombo; size++ {
		// Create a temporary slice to store current combination
		combo := make([]int, size)
		
		// Initialize the first combination
		for i := 0; i < size; i++ {
			combo[i] = i
		}
		
		for {
			// Generate current combination
			keyword := ""
			literal := ""
			ref := ""
			
			for _, idx := range combo {
				keyword += decorators[idx].getKeyword()
				if len(literal) > 0 {
					literal += ";"
				}
				literal += decorators[idx].getLiteral()
				ref += decorators[idx].getRef()
			}
			
			// Add to results
			result = append(result, struct {
				keyword string
				literal string
				ref     string
			}{keyword, literal, ref})
			
			// Generate next combination in lexicographic order
			// Find the rightmost element that can be incremented
			nextPos := size - 1
			for nextPos >= 0 && combo[nextPos] == n - size + nextPos {
				nextPos--
			}
			
			// If no such element exists, we're done
			if nextPos < 0 {
				break
			}
			
			// Increment the element and reset all elements to the right
			combo[nextPos]++
			for i := nextPos + 1; i < size; i++ {
				combo[i] = combo[i-1] + 1
			}
		}
	}
	
	return result
}

func LoadOmls() []*oml {
	colors := []*oml{
		newOml("Black", "30", "Blk"),
		newOml("Red", "31", "Red"),
		newOml("Green", "32", "Grn"),
		newOml("Yellow", "33", "Yel"),
		newOml("Blue", "34", "Blu"),
		newOml("Magenta", "35", "Mgt"),
		newOml("Purple", "35", "Prp"),
		newOml("Cyan", "36", "Cyn"),
		newOml("White", "37", "Wht"),
		newOml("BrightBlack", "90", "BrK"),
		newOml("BrightRed", "91", "BrR"),
		newOml("BrightGreen", "92", "BrG"),
		newOml("BrightYellow", "93", "BrY"),
		newOml("BrightBlue", "94", "BrB"),
		newOml("BrightMagenta", "95", "BrM"),
		newOml("BrightCyan", "96", "BrC"),
		newOml("BrightWhite", "97", "BrW"),
	}

	decorators := []*oml{
		newOml("Bold", "1", "b"),
		newOml("Dim", "2", "d"),
		newOml("Italic", "3", "i"),
		newOml("Underline", "4", "u"),
		newOml("SlowBlink", "5", "s"),
		newOml("RapidBlink", "6", "r"),
		newOml("Hidden", "8", "h"),
		newOml("Strike", "9", "t"),
	}

	backgrounds := []*oml{
		newOml("BBlack", "40", "BBlk"),
		newOml("BRed", "41", "BRed"),
		newOml("BGreen", "42", "BGrn"),
		newOml("BYellow", "43", "BYel"),
		newOml("BBlue", "44", "BBlu"),
		newOml("BMagenta", "45", "BMgt"),
		newOml("BPurple", "45", "BPrp"),
		newOml("BCyan", "46", "BCyn"),
		newOml("BWhite", "47", "BWht"),
		newOml("BBrightBlack", "100", "BBrK"),
		newOml("BBrightRed", "101", "BBrR"),
		newOml("BBrightGreen", "102", "BBrG"),
		newOml("BBrightYellow", "103", "BBrY"),
		newOml("BBrightBlue", "104", "BBrB"),
		newOml("BBrightMagenta", "105", "BBrM"),
		newOml("BBrightCyan", "106", "BBrC"),
		newOml("BBrightWhite", "107", "BBrW"),
	}
	
	// Generate all decorator combinations (up to 3)
	decoratorCombos := genDecoratorCombinations(decorators, 3)

	omls := make([]*oml, 0)

	// 1. Add reset option
	omls = append(omls, newOml("Reset", "\033[0m", "Z")) // Reset clears all
	
	// 2. Plain colors without decorators or backgrounds
	for _, color := range colors {
		omls = append(omls, newOml(
			color.getKeyword(),
			"\033[" + color.getLiteral() + "m",
			color.getRef(),
		))
	}
	
	// 3. Plain backgrounds without decorators or colors
	for _, bg := range backgrounds {
		omls = append(omls, newOml(
			bg.getKeyword(),
			"\033[" + bg.getLiteral() + "m",
			bg.getRef(),
		))
	}
	
	// 4. Color + Background combinations (no decorators)
	for _, color := range colors {
		for _, bg := range backgrounds {
			keyword := color.getKeyword() + bg.getKeyword() 
			literal := "\033[" + color.getLiteral() + ";" + bg.getLiteral() + "m"
			ref := color.getRef() + bg.getRef()
			
			omls = append(omls, newOml(keyword, literal, ref))
		}
	}

	// 5. Just decorators without colors or backgrounds
	for _, combo := range decoratorCombos {
		literal := "\033[" + combo.literal + "m"
		omls = append(omls, newOml(combo.keyword, literal, combo.ref))
	}

	// 6. Decorators + Color combinations
	for _, combo := range decoratorCombos {
		for _, color := range colors {
			keyword := combo.keyword + color.getKeyword()
			literal := "\033[" + combo.literal
			if len(combo.literal) > 0 {
				literal += ";"
			}
			literal += color.getLiteral() + "m"
			ref := combo.ref + color.getRef()
			
			omls = append(omls, newOml(keyword, literal, ref))
		}
	}
	
	// 7. Decorators + Background combinations
	for _, combo := range decoratorCombos {
		for _, bg := range backgrounds {
			keyword := combo.keyword + bg.getKeyword()
			literal := "\033[" + combo.literal
			if len(combo.literal) > 0 {
				literal += ";"
			}
			literal += bg.getLiteral() + "m"
			ref := combo.ref + bg.getRef()
			
			omls = append(omls, newOml(keyword, literal, ref))
		}
	}
	
	// 8. Decorators + Color + Background combinations
	for _, combo := range decoratorCombos {
		for _, color := range colors {
			for _, bg := range backgrounds {
				keyword := combo.keyword + color.getKeyword() + bg.getKeyword()
				literal := "\033[" + combo.literal
				if len(combo.literal) > 0 {
					literal += ";"
				}
				literal += color.getLiteral() + ";" + bg.getLiteral() + "m"
				ref := combo.ref + color.getRef() + bg.getRef()
				
				omls = append(omls, newOml(keyword, literal, ref))
			}
		}
	}

	return omls
}

func ReleaseOmls(omls []*oml) {
	omls = nil
}

func AddOml(existingOmls []*oml, keyword, literal, ref string) (omls []*oml, err error) {
	if keyword == "" {
		err = fmt.Errorf("keyword cannot be empty")
		return
	}
	if literal == "" {
		err = fmt.Errorf("literal cannot be empty")
		return
	}
	if ref == "" {
		err = fmt.Errorf("ref cannot be empty")
		return
	}

	if slices.ContainsFunc(existingOmls, func(o *oml) bool {
		return o.getKeyword() == keyword
	}) {
		err = fmt.Errorf("oml with keyword %s already exists", keyword)
		return
	}
	if slices.ContainsFunc(existingOmls, func(o *oml) bool {
		return o.getRef() == ref
	}) {
		err = fmt.Errorf("oml with ref %s already exists", ref)
		return
	}
	if slices.ContainsFunc(existingOmls, func(o *oml) bool {
		return o.getLiteral() == literal
	}) {
		err = fmt.Errorf("oml with literal %s already exists", literal)
		return
	}
	// TODO: Allow for multiple omls with the same literal by removing the literal exists error having oml-tracker track based on literals while keeping the keywords & refs to know what to replace(e.g. <Go>Content<lang>, if both have the same literal lang can close Go)

	oml := newOml(keyword, literal, ref)
	omls = append(omls, oml)
	return
}

func RemoveOml(omls []*oml, keyword string) []*oml {
	for i, oml := range omls {
		if oml.getKeyword() == keyword {
			omls = append(omls[:i], omls[i+1:]...)
			break
		}
	}
	return omls
}

func UpdateOml(omls []*oml, currentKeyword, newKeyword, newLiteral, newRef string) []*oml {
	for i, oml := range omls {
		if oml.getKeyword() == currentKeyword {
			omls[i] = newOml(newKeyword, newLiteral, newRef)
			break
		}
	}
	return omls
}