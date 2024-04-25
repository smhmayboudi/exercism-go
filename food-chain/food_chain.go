package foodchain

const TestVersion = 1

type Animal struct {
	name  string
	bline string
	toeat string
}

var foodchain = map[int]Animal{
	1: Animal{name: "fly", bline: "I don't know why she swallowed the fly. Perhaps she'll die.", toeat: ""},
	2: Animal{name: "spider", bline: "It wriggled and jiggled and tickled inside her.", toeat: "fly"},
	3: Animal{name: "bird", bline: "How absurd to swallow a bird!", toeat: "spider that wriggled and jiggled and tickled inside her"},
	4: Animal{name: "cat", bline: "Imagine that, to swallow a cat!", toeat: "bird"},
	5: Animal{name: "dog", bline: "What a hog, to swallow a dog!", toeat: "cat"},
	6: Animal{name: "goat", bline: "Just opened her throat and swallowed a goat!", toeat: "dog"},
	7: Animal{name: "cow", bline: "I don't know how she swallowed a cow!", toeat: "goat"},
	8: Animal{name: "horse", bline: "She's dead, of course!", toeat: ""},
}

func Verse(v int) string {
	var buffer string
	buffer = "I know an old lady who swallowed a " + foodchain[v].name + ".\n"
	buffer += foodchain[v].bline
	if foodchain[v].toeat == "" {
		return buffer
	}
	for i := v; i > 1; i-- {
		buffer += "\nShe swallowed the " + foodchain[i].name + " to catch the " + foodchain[i].toeat + "."
	}
	buffer += "\n" + foodchain[1].bline

	return buffer
}

func Verses(v, w int) string {
	var buffer string
	for i := v; i <= w; i++ {
		buffer += Verse(i)
		// spacing
		if i != w {
			buffer += "\n\n"
		}
	}

	return buffer
}

func Song() string {
	var buffer string
	for i := 1; i <= len(foodchain); i++ {
		buffer += Verse(i)
		// spacing
		if i != len(foodchain) {
			buffer += "\n\n"
		}
	}

	return buffer
}
