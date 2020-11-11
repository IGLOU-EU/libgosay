package gosaylib

type BubbleDef struct {
	Before    string
	OneLine   string
	FirstLine string
	Lines     string
	LastLine  string
	After     string
}

func (b *BubbleDef) Speak() {
	b.Before = " _ "
	b.OneLine = "< >"
	b.FirstLine = "/ \\"
	b.Lines = "| |"
	b.LastLine = "\\ /"
	b.After = " - "
}

func (b *BubbleDef) Think() {
	b.Before = " _ "
	b.OneLine = "( )"
	b.FirstLine = "( )"
	b.Lines = "( )"
	b.LastLine = "( )"
	b.After = " - "
}

func (b *BubbleDef) Whisper() {
	b.Before = ""
	b.OneLine = ""
	b.FirstLine = ""
	b.Lines = ""
	b.LastLine = ""
	b.After = ""
}

func (b *BubbleDef) Narrative() {
	b.Before = "==="
	b.OneLine = "| |"
	b.FirstLine = "| |"
	b.Lines = "| |"
	b.LastLine = "| |"
	b.After = "==="
}
