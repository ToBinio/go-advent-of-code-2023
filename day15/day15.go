package day15

import (
	"advent-of-code-2023/io"
	"strconv"
	"strings"
)

func Run() {
	line := io.ReadLines("resources/day15/input.txt")[0]

	var boxes []Box

	for i := 0; i < 256; i++ {
		boxes = append(boxes, Box{lenses: []Lens{}})
	}

	elements := strings.Split(line, ",")

	for _, element := range elements {
		if strings.HasSuffix(element, "-") {
			label, _ := strings.CutSuffix(element, "-")

			boxIndex := hash(label)
			boxes[boxIndex] = removeFromBox(boxes[boxIndex], label)
		} else {
			split := strings.Split(element, "=")

			label := split[0]
			focalLength, _ := strconv.Atoi(split[1])

			boxIndex := hash(label)

			boxes[boxIndex] = addLensToBox(boxes[boxIndex], Lens{label: label, focalLength: focalLength})
		}
	}

	println(computeFocusingPower(boxes))
}

func hash(text string) int {
	value := 0

	for _, char := range text {
		value += int(char)
		value *= 17
		value %= 256
	}

	return value
}

type Lens struct {
	label       string
	focalLength int
}

type Box struct {
	lenses []Lens
}

func removeFromBox(box Box, label string) Box {
	for i, lens := range box.lenses {
		if lens.label == label {
			box.lenses = append(box.lenses[:i], box.lenses[i+1:]...)
			break
		}
	}

	return box
}

func addLensToBox(box Box, lens Lens) Box {
	for i, otherLens := range box.lenses {
		if otherLens.label == lens.label {
			box.lenses[i] = lens
			return box
		}
	}

	box.lenses = append(box.lenses, lens)

	return box
}

func computeFocusingPower(boxes []Box) int {
	power := 0

	for boxIndex, box := range boxes {
		for lensIndex, lens := range box.lenses {
			power += (boxIndex + 1) * (lensIndex + 1) * lens.focalLength
		}
	}

	return power
}
