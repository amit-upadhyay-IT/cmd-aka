package enums

import (
	"errors"
	"log"
	"strings"
)

func GetEnumsMap(lines []string) (map[string][]string, error) {
	enumMap := make(map[string][]string)
	for index := 0; index < len(lines)-1; index++ {
		if strings.HasPrefix(lines[index], "const(") || strings.HasPrefix(lines[index], "const (") {
			enumName, enumValues, i, err := readConstBlock(lines, index)
			if err != nil {
				return nil, err
			}
			index = i
			enumMap[enumName] = enumValues
		}
	}
	return enumMap, nil
}

// read file contents
// returns:
// - name of the enum
// - values of the constants whose type is enum name
// - last row number where that enum is getting closed
func readConstBlock(lines []string, begIndex int) (string, []string, int, error) {
	var parenthesisCount int
	parenthesisCount = 1
	i := begIndex + 1
	enumName := ""
	var enumValues []string
	for parenthesisCount != 0 {
		line := lines[i]
		words := strings.Fields(line)
		// check for parenthesis
		if strings.Contains(line, "(") {
			parenthesisCount++
		}
		if strings.Contains(line, ")") {
			parenthesisCount--
		}
		// continue adding enum values
		// rules: 1) It should not be a comment, 2) If words has length more than 1, then second word should either
		// match the enumName or it should be a comment, if its neither of these then we assume its a declaration of
		// some other constant. 3) if words length is 1, we simply add it to enum values
		if len(words) > 0 && enumName != "" && parenthesisCount != 0 {
			if !strings.HasPrefix(words[0], "//") { // ignoring the comments
				if len(words) > 1 && (words[1] == enumName || strings.HasPrefix(words[1], "//")) {
					enumValues = append(enumValues, words[0])
				} else if len(words) == 1 {
					enumValues = append(enumValues, words[0])
				}
			}
		}
		// figure out enumName
		if len(words) > 1 && enumName == "" && parenthesisCount != 0{
			enumName = words[1]
			if enumName == "" {
				return "", nil, -1, errors.New("something is very bad")
			}
			// add enum value to the enumList
			enumValues = append(enumValues, words[0])
		}
		i++
	}
	if enumName == "" {
		log.Fatal("Shit!, cant figure out enum name")
	}
	return enumName, enumValues, i, nil
}

