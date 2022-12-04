package service

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"regexp"
	"strings"
)

func ReadFile(banner string) (map[rune][]string, error) {
	hash, err := CalculateHashSum(banner)
	if err != nil {
		return nil, err
	}
	if !Contains(hash) {
		return nil, errors.New("you have changed file auditor")
	}
	n := map[rune][]string{}
	readFile, err := os.Open(banner + ".txt")
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	character := []string{}
	keyRune := rune(31)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			n[keyRune] = character
			character = []string{}
			keyRune++
		}
		character = append(character, line)
	}
	n[keyRune] = character
	delete(n, rune(31))
	readFile.Close()
	return n, nil
}

func CalculateHashSum(banner string) (string, error) {
	f, err := os.Open(banner + ".txt")
	if err != nil {
		return "", err
	}
	defer f.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func Contains(fileName string) bool {
	for _, value := range hashSum {
		if value == fileName {
			return true
		}
	}
	return false
}

func sliceToString(strSlice []string) string {
	return strings.Join(strSlice, "")
}

func replace(str string) string {
	str = strings.ReplaceAll(str, "\\n", "\n")
	str = strings.ReplaceAll(str, "\r\n", "\n")
	str = strings.ReplaceAll(str, "\\t", "     ")
	re_up := regexp.MustCompile(`(\\)([a-zA-Z0-9!?']+)`)
	str = re_up.ReplaceAllString(str, "$2")
	return str
}
