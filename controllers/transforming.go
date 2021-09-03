package adds

import (
	"fmt"
	"strings"
)

func Encode(websitename string) string {
const ( // Kodowanie znaków specialnych
	end = "_38_"
	dot = "_46_"
	slash = "_47_"
	doubledot = "_58_"
	questionMark = "_63_"
	equal = "_61_"
	minus = "_45_"
	plus = "_43_"
)

	var url string = websitename // URL

	//Zmienianie znaków zpecialnych na nasze kodowanie
	url = strings.ReplaceAll(url, ".", dot)
	url = strings.ReplaceAll(url, "&", end)
	url = strings.ReplaceAll(url, "/", slash)
	url = strings.ReplaceAll(url, ":", doubledot)
	url = strings.ReplaceAll(url, "?", questionMark)
	url = strings.ReplaceAll(url, "=", equal)
	url = strings.ReplaceAll(url, "-", minus)
	url = strings.ReplaceAll(url, "+", plus)

	fmt.Println(url) // wyświetlanie tego

	return url
}