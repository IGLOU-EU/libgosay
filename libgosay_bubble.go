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
}

// Speak define a speaking bubble structure
func (b *BubbleDef) Speak() {
	b.Before = [3]rune{' ', '_', ' '}
	b.OneLine = [3]rune{'<', ' ', '>'}
	b.FirstLine = [3]rune{'/', ' ', '\\'}
	b.Lines = [3]rune{'|', ' ', '|'}
	b.LastLine = [3]rune{'\\', ' ', '/'}
	b.After = [3]rune{' ', '-', ' '}
}

// Think define a thinking bubble structure
func (b *BubbleDef) Think() {
	b.Before = [3]rune{' ', '_', ' '}
	b.OneLine = [3]rune{'(', ' ', ')'}
	b.FirstLine = [3]rune{'(', ' ', ')'}
	b.Lines = [3]rune{'(', ' ', ')'}
	b.LastLine = [3]rune{'(', ' ', ')'}
	b.After = [3]rune{' ', '-', ' '}
}

// Whisper define a whisp bubble structure
func (b *BubbleDef) Whisper() {
	b.Before = [3]rune{' ', '·', ' '}
	b.OneLine = [3]rune{':', ' ', ':'}
	b.FirstLine = [3]rune{':', ' ', ':'}
	b.Lines = [3]rune{':', ' ', ':'}
	b.LastLine = [3]rune{':', ' ', ':'}
	b.After = [3]rune{' ', '·', ' '}
}

// Narrative define a narrative text structure
func (b *BubbleDef) Narrative() {
	b.Before = [3]rune{'=', '=', '='}
	b.OneLine = [3]rune{'|', ' ', '|'}
	b.FirstLine = [3]rune{'|', ' ', '|'}
	b.Lines = [3]rune{'|', ' ', '|'}
	b.LastLine = [3]rune{'|', ' ', '|'}
	b.After = [3]rune{'=', '=', '='}
}
