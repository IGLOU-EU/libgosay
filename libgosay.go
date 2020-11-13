package libgosay

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

// Pimp is the main part of this lib
// You need to instance a Pimp struct for using it
type Pimp struct {
	Column int
	Said   string

	Body   string
	Eye    string
	Tongue string
	Tail   string

	Bubble BubbleDef
}

const gopher string = `
{{.Said}}
    {{.Tail}}   ˏ⋒___⋒ˎ
     {{.Tail}}  ▏ {{.Eye}}__{{.Eye}} ▕
        ▏  {{.Tongue}}  ▕
        ▏U    U▕
        ▏      ▕
        ˋ-U---U-ˊ
`

// Create make a Pimp struct, init default values and return it
func Create() Pimp {
	var p Pimp
	p.goSayInit()

	return p
}

// Say is for say some thing with the lib
func (p Pimp) Say(s string) (string, error) {
	var out bytes.Buffer

	if p.Column <= 0 {
		return "", fmt.Errorf("Column size can't be lower than 1")
	}

	if p.Body == "" {
		return "", fmt.Errorf("Body template can't be empty")
	}

	p.Said = bubbleMyStrings(
		splitStringToLen(s, p.Column),
		p.Bubble,
	)

	if p.Said == "" {
		p.Tail = " "
	}

	body, err := template.New("body").Parse(p.Body)
	if err != nil {
		return "", err
	}

	err = body.Execute(&out, p)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func (p *Pimp) goSayInit() {
	p.Column = 40

	p.Body = gopher
	p.Eye = "@"
	p.Tongue = "UU"
	p.Tail = "\\"

	p.Bubble.Speak()
}

func bubbleMyStrings(l []string, b BubbleDef) string {
	var lineLen int

	bubbleLen := len(l)
	bubbleLines := make([]string, bubbleLen+2)

	if bubbleLen > 1 {
		lineLen = maxLen(l)
	} else {
		lineLen = len(l[0])
	}

	if bubbleLen == 1 && lineLen == 0 {
		return ""
	}

	bubbleLines[0] = fmt.Sprintf("%c%s%c",
		b.Before[0],
		strings.Repeat(string(b.Before[1]), lineLen+2),
		b.Before[2],
	)

	if bubbleLen == 1 {
		bubbleLines[1] = fmt.Sprintf("%c%c%s%c%c",
			b.OneLine[0],
			b.OneLine[1],
			l[0],
			b.OneLine[1],
			b.OneLine[2],
		)
	} else {
		for i, v := range l {
			var decoLine [3]rune

			if i == 0 {
				decoLine = b.FirstLine
			} else if i == bubbleLen-1 {
				decoLine = b.LastLine
			} else {
				decoLine = b.Lines
			}

			bubbleLines[i+1] = fmt.Sprintf("%c%c%s%s%c%c",
				decoLine[0],
				decoLine[1],
				v,
				strings.Repeat(" ", lineLen-len(v)),
				decoLine[1],
				decoLine[2],
			)
		}
	}

	bubbleLines[bubbleLen+1] = fmt.Sprintf("%c%s%c",
		b.After[0],
		strings.Repeat(string(b.After[1]), lineLen+2),
		b.After[2],
	)

	return strings.Join(bubbleLines, "\n")
}

func maxLen(a []string) int {
	var l int

	for i := range a {
		alen := len(a[i])

		if alen > l {
			l = alen
		}
	}

	return l
}

func splitStringToLen(s string, llen int) []string {
	var lines []string
	lStart, lEnd := 0, 0

	if len(s) < llen {
		return []string{s}
	}

	for i, v := range s {
		if v == 9 || v == 10 || v == 32 {
			lEnd = i
		}

		if v == 10 || i-lStart >= llen {
			if lEnd <= lStart {
				lEnd = i
			}

			lines = append(lines, s[lStart:lEnd])

			if lEnd >= lStart {
				lStart = lEnd + 1
			} else {
				lStart = lEnd
			}
		}
	}

	return append(lines, s[lStart:])
}
