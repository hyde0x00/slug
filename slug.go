package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	cyrToLat = map[string]string{
		"а": "a", "б": "b", "в": "v", "г": "g", "д": "d", "е": "e", "ё": "e", "ж": "zh", "з": "z", "и": "i", "й": "y",
		"к": "k", "л": "l", "м": "m", "н": "n", "о": "o", "п": "p", "р": "r", "с": "s", "т": "t", "у": "u", "ф": "f",
		"х": "kh", "ц": "ts", "ч": "ch", "ш": "sh", "щ": "shch",
		"ъ": "", "ы": "y", "ь": "", "э": "e", "ю": "yu", "я": "ya",
	}
	keepCyrillic = flag.Bool("c", false, "keep cyrillic")
	dotsArg      = flag.Int("d", 0, "keep n dots at the end")
	sepArg       = flag.String("s", "-", "change spaces to separator")
	sep          = regexp.QuoteMeta(*sepArg)
)

func replaceDots(s string) string {
	numDots := strings.Count(s, ".") - *dotsArg
	if numDots < 0 {
		numDots = 0
	}
	s = strings.Replace(s, ".", sep, numDots)
	return s
}

func transliterate(s string) string {
	if *keepCyrillic == false {
		for from, to := range cyrToLat {
			s = strings.ReplaceAll(s, from, to)
		}
	}
	return s
}

func slugify(s string) string {
	s = replaceDots(s)

	re := regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9./_\t]+`)
	s = re.ReplaceAllString(s, sep)
	s = strings.Trim(s, sep)
	s = strings.ReplaceAll(s, sep+".", ".")

	return s
}

func printResult(s string) {
	s = strings.ToLower(s)
	s = transliterate(s)
	s = slugify(s)
	fmt.Println(s)
}

func main() {
	flag.Parse()

	input, _ := os.Stdin.Stat()

	if input.Mode()&os.ModeCharDevice == 0 {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			printResult(s.Text())
		}
	}
	for _, arg := range flag.Args() {
		printResult(arg)
	}
}
