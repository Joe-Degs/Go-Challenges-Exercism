// Package bob helps bob a teenager to respond to remarks.
package bob

// bob represents the teenager bob.
type Bob struct {
	remark   string // bob tries to provide a response to remark.
	response string // bob's response to the remark.
}

// remarkIsCaps checks if all letters in remark is in uppercase.
func (b *Bob) remarkIsCaps() bool {
	// return false if remark contains no english alphabet
	if b.remarkHasNoLetters() {
		return false
	}

	for _, r := range b.remark {
		// return false if char is ascii letter and
		// is not an uppercase letter.
		if isLetter(r) && !(r >= 'A' && r <= 'Z') {
			return false
		}
	}
	return true
}

// remarkHasNoLetters checks to see if remark contains no letters.
func (b *Bob) remarkHasNoLetters() bool {
	for _, r := range b.remark {
		if isLetter(r) {
			return false
		}
	}
	return true
}

// remarkEndsWithWhiteSpace scans a remark that contains letters
// to see if it ends with a whitespace character(s).
func (b *Bob) remarkEndsWithWhiteSpace() bool {
	if !b.remarkHasNoLetters() && isWhiteSpace([]rune(b.remark)[len(b.remark)-1]) {
		return true
	}
	return false
}

// removeWhiteSpaceFromRemarks removes whitespace from remark
// from the end of the string.
// should only be called if remarkEndsWithWhiteSpace returns true.
func (b *Bob) removeWhiteSpaceFromRemark() {
	r := []rune(b.remark)
	for i := len(r) - 1; i >= 0; i-- {
		if !isWhiteSpace(r[i]) {
			b.remark = string(r[0 : i+1])
			break
		}
	}
}

// remarkEndsWithQueMark checks to see if remark ends with question mark.
func (b *Bob) remarkEndsWithQueMark() bool {
	if []rune(b.remark)[len(b.remark)-1] == '?' {
		return true
	}
	return false
}

// remarkIsWhiteSpace returns true if remark is only whitespace character(s).
func (b *Bob) remarkIsWhiteSpace() bool {
	for _, r := range b.remark {
		if !isWhiteSpace(r) {
			return false
		}
	}
	return true
}

// respond provides a response to a remark.
func (b *Bob) respond() (response string) {

	// remove trailing whitespace from end of remark.
	if b.remarkEndsWithWhiteSpace() {
		b.removeWhiteSpaceFromRemark()
	}

	if b.remarkIsWhiteSpace() {
		b.response = "Fine. Be that way!"
		return b.response
	} else if b.remarkIsCaps() {
		if b.remarkEndsWithQueMark() {
			b.response = "Calm down, I know what I'm doing!"
			return b.response
		}
		b.response = "Whoa, chill out!"
		return b.response
	} else if b.remarkEndsWithQueMark() {
		b.response = "Sure."
		return b.response
	} else {
		b.response = "Whatever."
		return b.response
	}
}

// -- Helper Functions

// isWhiteSpace checks whether a given character
// is a whitespace character.
func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// isLetter checks if a character part of the english alphabet.
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// Hey returns bob's response to a remark.
func Hey(remark string) string {
	bob := &Bob{remark: remark}
	return bob.respond()
}
