package util

import (
	"math/rand"
	"strings"
	"time"
)

// set the seed value for the random generator
func init() {
	// As rand.Seed() expect an int64 as input,
	// convert the time to unix nano before passing it to the function.
	rand.Seed(time.Now().UnixNano())
}

// takes 2 int64 numbers: min and max as input.
// And it returns a random int64 number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generate a random string of n characters
const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomTitle() string {
	return RandomString(6)
}

func generateRandomRune(size int, start, end int64) string {
	randRune := make([]rune, size)
	for i := range randRune {
		randRune[i] = rune(RandomInt(start, end-1))
	}
	return string(randRune)
}

func RandomJapaneseTitle() string {
	japaneseHiragana := []int64{12353, 12435}

	return generateRandomRune(10, japaneseHiragana[0], japaneseHiragana[1])
}

func RandomDay() int64 {
	return RandomInt(1, 29)
}

func RandomMonth() int64 {
	return RandomInt(1, 12)
}

func RandomSource() string {
	sources := []string{"Manga", "Light novel", "Web manga", "Original", "Visual novel"}
	n := len(sources)
	return sources[rand.Intn(n)]
}

func RandomGenres() []string {
	genres := []string{
		"Sci-Fi", "Mystery", "Horror", "Psychological", "Thriller",
		"Action", "Adventure", "Drama", "Fantasy", "Comedy",
		"Romance", "School", "Seinen", "Slice of Life", "Supernatural",
		"Demons", "Historical", "Samurai", "Shounen", "Game",
		"Harem", "Mecha", "Sports", "Magic", "Dementia", "Super Power",
	}
	rand.Shuffle(len(genres), func(i, j int) { genres[i], genres[j] = genres[j], genres[i] })

	return genres[:5]
}

func RandomDescription() string {
	return "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
}

func RandomSeason() string {
	seasons := []string{
		"Sprint", "Summer", "Fall", "Winter",
	}
	n := len(seasons)
	return seasons[rand.Intn(n)]
}

func RandomYear() int64 {
	return RandomInt(2016, 2021)
}

func RandomNumEpisodes() int64 {
	return RandomInt(12, 24)
}

// TODO: randomize rest of the columns
