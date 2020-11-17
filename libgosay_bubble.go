package libgosay

// BubbleDef is for defined the bubble structure
// That cane be done manually or using a pre configured structure
type BubbleDef struct {
	Before    [3]rune
	OneLine   [3]rune
	FirstLine [3]rune
	Lines     [3]rune
	LastLine  [3]rune
	After     [3]rune
	Tail      rune
}

// Speak define a speaking bubble structure
func (b *BubbleDef) Speak() {
	b.Before = [3]rune{' ', '_', ' '}
	b.OneLine = [3]rune{'<', ' ', '>'}
	b.FirstLine = [3]rune{'/', ' ', '\\'}
	b.Lines = [3]rune{'|', ' ', '|'}
	b.LastLine = [3]rune{'\\', ' ', '/'}
	b.After = [3]rune{' ', '-', ' '}
	b.Tail = '\\'
}

// Think define a thinking bubble structure
func (b *BubbleDef) Think() {
	b.Before = [3]rune{' ', '_', ' '}
	b.OneLine = [3]rune{'(', ' ', ')'}
	b.FirstLine = [3]rune{'(', ' ', ')'}
	b.Lines = [3]rune{'(', ' ', ')'}
	b.LastLine = [3]rune{'(', ' ', ')'}
	b.After = [3]rune{' ', '-', ' '}
	b.Tail = 'o'
}

// ThinkUTF8 define a thinking bubble structure but in UTF-8
func (b *BubbleDef) ThinkUTF8() {
	b.Before = [3]rune{' ', '⁀', ' '}
	b.OneLine = [3]rune{'(', ' ', ')'}
	b.FirstLine = [3]rune{'(', ' ', ')'}
	b.Lines = [3]rune{'(', ' ', ')'}
	b.LastLine = [3]rune{'(', ' ', ')'}
	b.After = [3]rune{' ', '‿', ' '}
	b.Tail = 'o'
}

// Whisper define a whisp bubble structure
func (b *BubbleDef) Whisper() {
	b.Before = [3]rune{' ', '.', ' '}
	b.OneLine = [3]rune{':', ' ', ':'}
	b.FirstLine = [3]rune{':', ' ', ':'}
	b.Lines = [3]rune{':', ' ', ':'}
	b.LastLine = [3]rune{':', ' ', ':'}
	b.After = [3]rune{' ', '.', ' '}
	b.Tail = '*'
}

// WhisperUTF8 define a whisp bubble structure but in UTF-8
func (b *BubbleDef) WhisperUTF8() {
	b.Before = [3]rune{'◜', '╌', '◝'}
	b.OneLine = [3]rune{'╎', ' ', '╎'}
	b.FirstLine = [3]rune{'╎', ' ', '╎'}
	b.Lines = [3]rune{'╎', ' ', '╎'}
	b.LastLine = [3]rune{'╎', ' ', '╎'}
	b.After = [3]rune{'◟', '╌', '◞'}
	b.Tail = '⋱'
}

// Narrative define a narrative text structure
func (b *BubbleDef) Narrative() {
	b.Before = [3]rune{'=', '=', '='}
	b.OneLine = [3]rune{'|', ' ', '|'}
	b.FirstLine = [3]rune{'|', ' ', '|'}
	b.Lines = [3]rune{'|', ' ', '|'}
	b.LastLine = [3]rune{'|', ' ', '|'}
	b.After = [3]rune{'=', '=', '='}
	b.Tail = ' '
}

// BigBoxUTF8 define a narrative text structure with a big box
func (b *BubbleDef) BigBoxUTF8() {
	b.Before = [3]rune{'┏', '━', '┓'}
	b.OneLine = [3]rune{'┃', ' ', '┃'}
	b.FirstLine = [3]rune{'┃', ' ', '┃'}
	b.Lines = [3]rune{'┃', ' ', '┃'}
	b.LastLine = [3]rune{'┃', ' ', '┃'}
	b.After = [3]rune{'┗', '━', '┛'}
	b.Tail = ' '
}

// AsianBoxUTF8 define a narrative text structure with asian style
func (b *BubbleDef) AsianBoxUTF8() {
	b.Before = [3]rune{'㇛', '㇐', '㇕'}
	b.OneLine = [3]rune{'㇑', ' ', '㇑'}
	b.FirstLine = [3]rune{'㇑', ' ', '㇑'}
	b.Lines = [3]rune{'㇑', ' ', '㇑'}
	b.LastLine = [3]rune{'㇑', ' ', '㇑'}
	b.After = [3]rune{'㇗', '㇐', '㇘'}
	b.Tail = ' '
}
