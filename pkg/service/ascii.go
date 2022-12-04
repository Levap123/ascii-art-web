package service

import (
	"fmt"
	"strings"
)

var hashSum = []string{"86d9947457f6a41a18cb98427e314ff8", "a49d5fcb0d5c59b2e77674aa3ab8bbb1", "a51f800619146db0c42d26db3114c99f"}

const bannersPath = "src/banners/"

type AsciiService struct{}

type Ascii interface {
	GenerateWords(target, banner string) (string, error)
}

func newAsciiService() *AsciiService {
	return &AsciiService{}
}

func (s *AsciiService) GenerateWords(target, banner string) (string, error) {
	charsMap, err := ReadFile(bannersPath + banner)
	target = replace(target)
	if err != nil {
		return "", err
	}
	arr := strings.Split(target, "\n")
	words := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) != 0 {
			for index := 1; index <= 8; index++ {
				for _, char := range arr[i] {
					if val, ok := charsMap[char]; ok {
						words[i] += val[index]
					} else {
						return "", fmt.Errorf("character in argument is not allowed\n")
					}
				}
				words[i] += "\n"
			}
		} else {
			if i != 0 {
				words[i] += "\n"
			}
		}
	}

	return sliceToString(words), nil
}
