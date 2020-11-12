package libgosay

type BubbleDef struct {
	Before    [3]rune
	OneLine   [3]rune
	FirstLine [3]rune
	Lines     [3]rune
	LastLine  [3]rune
	After     [3]rune
}

func (b *BubbleDef) Speak() {
	b.Before = [3]rune{' ', '_', ' '}
	b.OneLine = [3]rune{'<', ' ', '>'}
	b.FirstLine = [3]rune{'/', ' ', '\\'}
	b.Lines = [3]rune{'|', ' ', '|'}
	b.LastLine = [3]rune{'\\', ' ', '/'}
	b.After = [3]rune{' ', '-', ' '}
}

func (b *BubbleDef) Think() {
	b.Before = [3]rune{' ', '_', ' '}
	b.OneLine = [3]rune{'(', ' ', ')'}
	b.FirstLine = [3]rune{'(', ' ', ')'}
	b.Lines = [3]rune{'(', ' ', ')'}
	b.LastLine = [3]rune{'(', ' ', ')'}
	b.After = [3]rune{' ', '-', ' '}
}

func (b *BubbleDef) Whisper() {
	b.Before = [3]rune{' ', '·', ' '}
	b.OneLine = [3]rune{':', ' ', ':'}
	b.FirstLine = [3]rune{':', ' ', ':'}
	b.Lines = [3]rune{':', ' ', ':'}
	b.LastLine = [3]rune{':', ' ', ':'}
	b.After = [3]rune{' ', '·', ' '}
}

func (b *BubbleDef) Narrative() {
	b.Before = [3]rune{'=', '=', '='}
	b.OneLine = [3]rune{'|', ' ', '|'}
	b.FirstLine = [3]rune{'|', ' ', '|'}
	b.Lines = [3]rune{'|', ' ', '|'}
	b.LastLine = [3]rune{'|', ' ', '|'}
	b.After = [3]rune{'=', '=', '='}
}
