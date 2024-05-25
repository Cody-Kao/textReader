package cmd

import (
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"

	"github.com/AlecAivazis/survey/v2"
)

func Read() {
	var option string
	prompt := &survey.Select{
		Message: "characters to count(or choose Exit to terminate the program)",
		Options: []string{"All", "Only English Letter", "Only Chinese Character", "Only Special Character", "Only Number", "Exit"},
	}

	err := survey.AskOne(prompt, &option)
	if err != nil {
		fmt.Println("Error when choosing options")
		os.Exit(2) // status code for invalid option
	}

	if option == "Exit" {
		fmt.Println("Terminate the program")
		return
	}

	fmt.Println("chosen option:", option)

	// open file
	fmt.Println("path", Filepath)
	file, err := os.Open(Filepath)
	if err != nil {
		fmt.Println("invalid file path")
		os.Exit(127) // status code for path not finding
	}
	defer file.Close()

	// start reading
	var count int
	buffer := make([]byte, 1024)

	for {
		numOfBytes, err := file.Read(buffer)
		for i := 0; i < numOfBytes; {
			// 因為英文字佔1byte而中文字會佔3個byte，所以要先轉換為rune再看是否為英文或中文
			// index i應該要跟著字元數而增加，所以加上size(1代表一個英文字，3就是代表一個中文字)
			ch, size := utf8.DecodeRune(buffer[i:numOfBytes]) // byte => rune
			i += size

			if option == "Only English Letter" && isEnglishLetter(ch) {
				count++
			}
			if option == "Only Chinese Character" && isChineseCharacter(ch) {
				count++
			}
			if option == "Only Special Character" && isSpecialCharacter(ch) {
				count++
			}
			if option == "Only Number" && isNumber(ch) {
				count++
			}
			if option == "All" {
				count++
			}
		}
		// 檔案讀完會報的錯誤
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	fmt.Println("total characters are:", count)
}

func isEnglishLetter(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

func isChineseCharacter(ch rune) bool {
	return unicode.Is(unicode.Han, ch)
}

// isSpecialCharacter checks if a rune is a special character
func isSpecialCharacter(ch rune) bool {
	return unicode.IsPunct(ch) || unicode.IsSymbol(ch) || unicode.IsSpace(ch)
}

func isNumber(ch rune) bool {
	return unicode.IsDigit(ch)
}
