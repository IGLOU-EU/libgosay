package libgosay

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

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
    {{.Tail}}   ˏ⋒____⋒ˎ
     {{.Tail}}  ▏ {{.Eye}}__{{.Eye}} ▕
        ▏  {{.Tongue}}  ▕
        ▏U    U▕
        ▏      ▕
        ˋ-U--U-ˊ
`

func (p *Pimp) Default() {
	p.Column = 40

	p.Body = gopher
	p.Eye = "@"
	p.Tongue = "UU"
	p.Tail = "\\"

	p.Bubble.Speak()
}

func (p Pimp) Say(s string) (string, error) {
	var out bytes.Buffer

	if p.Column <= 0 {
		return "", fmt.Errorf("Column size can't be lower than 1")
	}

	p.Said = bubbleMyStrings(
		splitStringToLen(s, p.Column),
		p.Bubble,
	)

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

func bubbleMyStrings(l []string, b BubbleDef) string {
	var lineLen int

	bubbleLen := len(l)
	bubbleLines := make([]string, bubbleLen+2)

	if bubbleLen > 1 {
		lineLen = maxLen(l)
	} else {
		lineLen = len(l[0])
	}

	bubbleLines[0] = fmt.Sprintf("%s%s%s",
		b.Before[:1],
		strings.Repeat(b.Before[1:2], lineLen+2),
		b.Before[2:],
	)

	if bubbleLen == 1 {
		bubbleLines[1] = fmt.Sprintf("%s%s%s%s%s",
			b.OneLine[:1],
			b.OneLine[1:2],
			l[0],
			b.OneLine[1:2],
			b.OneLine[2:],
		)
	} else {
		for i, v := range l {
			var decoLine string
			var spacer string

			if i == 0 {
				decoLine = b.FirstLine
			} else if i == bubbleLen-1 {
				decoLine = b.LastLine
			} else {
				decoLine = b.Lines
			}

			spacer = strings.Repeat(" ", lineLen-len(v))

			bubbleLines[i+1] = fmt.Sprintf("%s%s%s%s%s%s",
				decoLine[:1],
				decoLine[1:2],
				v,
				spacer,
				decoLine[1:2],
				decoLine[2:],
			)
		}
	}

	bubbleLines[bubbleLen+1] = fmt.Sprintf("%s%s%s",
		b.After[:1],
		strings.Repeat(b.After[1:2], lineLen+2),
		b.After[2:],
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
