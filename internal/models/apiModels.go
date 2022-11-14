package models

import "github.com/wenlng/go-captcha/captcha"

type SendData struct {
	PassWord string
	Picture  string
}

func CreatorCaptcha() (string, string, error) {

	s1 := captcha.Size{
		Width:  1,
		Height: 1,
	}

	s2 := captcha.Size{
		Width:  199,
		Height: 150,
	}

	alfabetInput := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "a", "B", "b", "C", "c", "D", "d", "E", "e", "F", "f", "G", "g", "H", "h", "I", "i", "J", "j", "K", "k", "L", "l", "M", "m", "N", "n", "O", "o", "P", "p", "Q", "q", "R", "r", "S", "s", "T", "t", "U", "u", "V", "v", "W", "w", "X", "x", "Y", "y", "Z", "z"}
	//	alfabetInput := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	capt := captcha.GetCaptcha()
	capt.SetRangChars(alfabetInput)
	capt.SetRangCheckTextLen(captcha.RangeVal{Min: 4, Max: 6})
	capt.SetImageQuality(99)
	dots, b64, tb64, key, err := capt.GenerateWithSize(s1, s2) // capt.Generate()

	var keyStr string
	for _, t1 := range dots {
		keyStr = keyStr + t1.Text
	}

	_ = b64
	_ = key

	//	fmt.Println(keyStr)
	//	fmt.Println(`++++++++`)
	//	fmt.Println(tb64)

	return keyStr, tb64, err
}
