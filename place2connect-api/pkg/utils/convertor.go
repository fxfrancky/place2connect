package utils

import (
	"log"
	"os"
	"strconv"
)

func UIntToString(ui uint) string {
	return strconv.FormatUint(uint64(ui), 10)
}

func StringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		// return 0, errors.New("couldn't convert a string to int")
		log.Println("************** Printing this ", s)
		return 0, nil
	}
	return num, nil
}
func StringToBool(s string) (bool, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		// return 0, errors.New("couldn't convert a string to int")
		log.Println("************** Printing this ", s)
		return false, nil
	}
	return b, nil
}

func FileExists(filename string) (bool, error) {
	fullPath := "./images/" + filename
	log.Println("The full path is")
	info, err := os.Stat(fullPath)
	// os.OpenFile(name string, flag int, perm os.FileMode)
	// info.
	if os.IsNotExist(err) {
		log.Printf("The file %s doesnt exist ", fullPath)
		return false, err
	}
	return !info.IsDir(), nil
}
