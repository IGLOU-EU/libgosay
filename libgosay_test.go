package libgosay

import (
	"fmt"
	"testing"
)

func tError(r bool, s string, t *testing.T) {
	if r {
		t.Errorf("%s", s)
	}
}

func TestStrfunc(t *testing.T) {
	ss := splitStringToLen("Go for the eyes, Boo, go for the eyes!", 10)
	tError(len(ss) != 4, fmt.Sprint("\nStringSplit error:\nExpect-4 give-", len(ss), "\n", ss), t)

	ml := maxLen(ss)
	tError(ml != 10, fmt.Sprint("\nMax array len error:\nExpect-10 give-", ml), t)
}

func TestSay(t *testing.T) {
	var g Pimp
	g.Default()

	se := 199
	s, e := g.Say("Heya, it's me Imoen")
	tError(e != nil, fmt.Sprint("DefaultSay error:", e), t)
	tError(len(s) != se, fmt.Sprint("\nDefaultSay render len don't match\nExpected-", se, " Rended-", len(s), s), t)

	se = 479
	s, e = g.Say("I've seen things you people wouldn't believe. Attack ships on fire off the shoulder of Orion. I watched C-beams glitter in the dark near the Tannhauser gate. All those moments will be lost in time, like tears in rain. Time to die.")
	tError(e != nil, fmt.Sprint("DefaultSay error:", e), t)
	tError(len(s) != se, fmt.Sprint("\nDefaultSay render len don't match\nExpected-", se, " Rended-", len(s), s), t)
}

func TestBubble(t *testing.T) {
	var b BubbleDef
	s := splitStringToLen("Les méchants, je les démolis. Tout ce qui brille, je me le garde. Devant moi, mes ennemis détalent. Je suis née pour gagner !", 40)

	er := ` __________________________________________ 
/ Les méchants, je les démolis. Tout ce  \
| qui brille, je me le garde. Devant moi,  |
| mes ennemis détalent. Je suis née pour |
\ gagner !                                 /
 ------------------------------------------ `
	b.Speak()
	gr := bubbleMyStrings(s, b)
	tError(gr != er, fmt.Sprint("Bad bubble formating:\n", gr), t)

	s = splitStringToLen("I am the law!", 15)

	er = ` _______________ 
< I am the law! >
 --------------- `
	b.Speak()
	gr = bubbleMyStrings(s, b)
	tError(gr != er, fmt.Sprint("Bad bubble formating:\n", gr), t)

	er = ` _______________ 
( I am the law! )
 --------------- `
	b.Think()
	gr = bubbleMyStrings(s, b)
	tError(gr != er, fmt.Sprint("Bad bubble formating:\n", gr), t)

	er = ` ··············· 
: I am the law! :
 ··············· `
	b.Whisper()
	gr = bubbleMyStrings(s, b)
	tError(gr != er, fmt.Sprint("Bad bubble formating:\n", gr), t)

	er = `=================
| I am the law! |
=================`
	b.Narrative()
	gr = bubbleMyStrings(s, b)
	tError(gr != er, fmt.Sprint("Bad bubble formating:\n", gr), t)
}
