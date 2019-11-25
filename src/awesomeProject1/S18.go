package main
//strings包
import (
	"fmt"
	"strings"
)

func urlProcess(url string) string {
	result := strings.HasPrefix(url, "hhtp://")
	if !result {
		url = fmt.Sprintf("http://%s",url)
	}
	return url
}

func pathProcess(path string) string {
	result := strings.HasSuffix(path, "hhtp://")
	if !result {
		path = fmt.Sprintf("%s/",path)
	}
	return path
}

func main() {
	//var(
	//	url string
	//	path string
	//)
	//fmt.Scanf("%s%s",&url , &path)
	//url = urlProcess(url)
	//path = pathProcess(path)
	//fmt.Printf("url = %s path = %s",url , path)

	str := "  Hello world abc  \n"
	result := strings.Replace(str,"world","you",1)
	fmt.Printf("result = \n",result)

	count := strings.Count(str,"l")
	fmt.Printf("count = \n",count)

	result1 := strings.Repeat(str,3)
	fmt.Printf("result1 = \n",result1)

	//ToLower ToUpper  略

}