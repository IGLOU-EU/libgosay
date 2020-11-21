package libgosay

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

func tError(r bool, s string, t *testing.T) {
	if r {
		t.Errorf("%s", s)
	}
}

func TestErrorfuc(t *testing.T) {
	var p Pimp
	var err error

	_, err = p.Say("")
	tError(err == nil, "\nIf Say() is call with an Column size zero, we expect to catch an err", t)
	p.Column = 10

	_, err = p.Say("")
	tError(err == nil, "\nIf Say() is call with an empty p.Body, we expect to catch an err", t)

	p.setEyes()
	tError(p.EyeL != "" || p.EyeR != "", "\nWe expect EyeL and EyeR are empty", t)

	p.Eyes = "0Oo"
	p.setEyes()
	tError(p.EyeL != "0" || p.EyeR != "O", fmt.Sprintf("\nWe expect to EyeL == 0 and EyeR == O :\np.Eyes = %s\np.EyeL = %s && p.EyeL = %s", p.Eyes, p.EyeL, p.EyeR), t)

	p.Eyes = "--"
	p.setEyes()
	tError(p.EyeL != "0" || p.EyeR != "O", fmt.Sprintf("\nWe expect Eye* not change to -- :\np.Eyes = %s\np.EyeL = %s && p.EyeL = %s", p.Eyes, p.EyeL, p.EyeR), t)

	p.Tail = "o"
	p.setTail()
	tError(p.Tail != "o", fmt.Sprintf("We expect p.Tail == o but p.Tail == %s", p.Tail), t)
}

func TestStrfunc(t *testing.T) {
	ss := splitStringToLen("Go for the eyes, Boo, go for the eyes!", 10)
	tError(len(ss) != 4, fmt.Sprint("\nStringSplit error:\nExpect-4 give-", len(ss), "\n", ss), t)

	ml := maxLen(ss)
	tError(ml != 10, fmt.Sprint("\nMax array len error:\nExpect-10 give-", ml), t)

	i := spaceBefore("toto", "___toto")
	tError(i != 3, fmt.Sprint("\nSpace before object:\nExpect-3 give-", i), t)
}

func TestSay(t *testing.T) {
	var g = Create()

	se := 128
	s, e := g.Say("")
	tError(e != nil, fmt.Sprint("DefaultSay error:", e), t)
	tError(len(s) != se, fmt.Sprint("\nDefaultSay render len don't match\nExpected-", se, " Rended-", len(s), s), t)

	se = 199
	s, e = g.Say("Heya, it's me Imoen")
	tError(e != nil, fmt.Sprint("DefaultSay error:", e), t)
	tError(len(s) != se, fmt.Sprint("\nDefaultSay render len don't match\nExpected-", se, " Rended-", len(s), s), t)

	se = 479
	s, e = g.Say("I've seen things you people wouldn't believe. Attack ships on fire off the shoulder of Orion. I watched C-beams glitter in the dark near the Tannhauser gate. All those moments will be lost in time, like tears in rain. Time to die.")
	tError(e != nil, fmt.Sprint("DefaultSay error:", e), t)
	tError(len(s) != se, fmt.Sprint("\nDefaultSay render len don't match\nExpected-", se, " Rended-", len(s), s), t)
}

func TestCustomSay(t *testing.T) {
	var g = Create()
	g.Body = `{{- /* From https://github.com/paulkaefer/cowsay-files.git */ -}}
{{- /* An owl */ -}}
 {{.Said}}
		 {{.Tail}}
		  {{.Tail}}
		   ___
		  (o o)
		 (  V  )
		/--m-m-`

	expect := `  _____________________ 
 < Heya, it's me Imoen >
  --------------------- 
		 \
		  \
		   ___
		  (o o)
		 (  V  )
		/--m-m-`

	out, err := g.Say("Heya, it's me Imoen")
	tError(err != nil, fmt.Sprint("CustomSay error:", err), t)
	tError(out != expect, fmt.Sprint("Bad output formating:\n", out), t)
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
	gr := bubbleMyStrings(s, b, 0)
	tError(gr != er, fmt.Sprint("Bad bubble formating:\n", gr), t)

	s = splitStringToLen("I am the law!", 15)

	er = ` _______________ 
< I am the law! >
 --------------- `
	b.Speak()
	gr = bubbleMyStrings(s, b, 0)
	tError(gr != er, fmt.Sprint("Bad bubble formating:\n", gr), t)

	er = ` _______________ 
( I am the law! )
 --------------- `
	b.Think()
	gr = bubbleMyStrings(s, b, 0)
	tError(gr != er, fmt.Sprint("Bad bubble formating:\n", gr), t)

	er = ` ............... 
: I am the law! :
 ............... `
	b.Whisper()
	gr = bubbleMyStrings(s, b, 0)
	tError(gr != er, fmt.Sprint("Bad bubble formating:\n", gr), t)

	er = `=================
| I am the law! |
=================`
	b.Narrative()
	gr = bubbleMyStrings(s, b, 0)
	tError(gr != er, fmt.Sprint("Bad bubble formating:\n", gr), t)

	er = `     =================
     | I am the law! |
     =================`
	b.Narrative()
	gr = bubbleMyStrings(s, b, 5)
	tError(gr != er, fmt.Sprint("Bad spacer size:\n", gr), t)
}

var gs Pimp
var bfInt int
var bfStr string
var bfStrt []string

const gSay = "Les méchants, je les démolis. Tout ce qui brille, je me le garde. Devant moi, mes ennemis détalent. Je suis née pour gagner !"

func BenchmarkCreate(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gs = Create()
	}

	_ = gs
}

func BenchmarkSay(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gs.Say(gSay)
	}
}

func BenchmarkBms(b *testing.B) {
	var s string
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bubbleMyStrings([]string{gSay}, gs.Bubble, 0)
	}

	bfStr = s
}

func BenchmarkSstl(b *testing.B) {
	var s []string
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s = splitStringToLen(gSay, gs.Column)
	}

	bfStrt = s
}

func BenchmarkMl(b *testing.B) {
	var m int
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m = maxLen([]string{gSay})
	}

	bfInt = m
}

func BenchmarkBubble(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gs.Bubble.Speak()
		gs.Bubble.Think()
		gs.Bubble.ThinkUTF8()
		gs.Bubble.Whisper()
		gs.Bubble.WhisperUTF8()
		gs.Bubble.Narrative()
		gs.Bubble.BigBoxUTF8()
		gs.Bubble.AsianBoxUTF8()
	}
}

func BenchmarkTpl(b *testing.B) {
	var s string
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var out bytes.Buffer
		body, _ := template.New("body").Parse(gs.Body)
		body.Execute(&out, gs)
		s = out.String()
	}

	bfStr = s
}
