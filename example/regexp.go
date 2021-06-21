package main

import "regexp"

func main() {
	path := "api-83248239498043-bjnksdf.appstore.com"
	var reg = regexp.MustCompile(`^api-(\d+)(-)([\w-]+)(\.appstore\.com)$`)
	match := reg.FindStringSubmatch(path)
	result := make(map[string]string)
	for i, name := range match {
		if i != 0 && name != "" { // 第一个分组为空（也就是整个匹配）
			result[name] = match[i]
			println(result[name] + ": " + match[i])
		}
	}
}
