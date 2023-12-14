package day13

import "advent-of-code-2023/io"

func Run() {
	lines := io.ReadLines("resources/day13/input.txt")

	var notes []Note

	currentNote := Note{data: [][]rune{}}
	for _, line := range lines {
		if line == "" {
			notes = append(notes, currentNote)
			currentNote = Note{data: [][]rune{}}
			continue
		}

		var charLine []rune

		for _, char := range line {
			charLine = append(charLine, char)
		}

		currentNote.data = append(currentNote.data, charLine)
	}
	notes = append(notes, currentNote)

	sum := 0

	for _, note := range notes {
		sum += getMirrorValue(note)
	}

	println(sum)
}

type Note struct {
	data [][]rune
}

func getMirrorValue(note Note) int {
	for i := 0; i < len(note.data)-1; i++ {
		if isMirroredHorizontal(note, i) == 1 {
			return (i + 1) * 100
		}
	}

	for i := 0; i < len(note.data[0])-1; i++ {
		if isMirroredVertically(note, i) == 1 {
			return i + 1
		}
	}

	return 0
}

func isMirroredVertically(note Note, index int) int {
	errors := 0

	for i := 0; i <= index; i++ {
		left := index - i
		right := index + i + 1

		if right >= len(note.data[0]) {
			break
		}

		for _, line := range note.data {
			if line[left] != line[right] {
				errors++
			}
		}
	}

	return errors
}

func isMirroredHorizontal(note Note, index int) int {
	errors := 0

	for i := 0; i <= index; i++ {
		left := index - i
		right := index + i + 1

		if right >= len(note.data) {
			break
		}

		for column := 0; column < len(note.data[0]); column++ {
			if note.data[left][column] != note.data[right][column] {
				errors++
			}
		}
	}

	return errors
}
