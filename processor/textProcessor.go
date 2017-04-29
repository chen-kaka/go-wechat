package processor

func ProcessText(msg string) string {
	personalBlog := "<a href='http://kaka.xy-kaka.cn'>个人博客</a>"
	technologyBlog := "<a href='https://chen-kaka.github.io'>个人技术博客</a>"
	fundPage := "<a href='http://fund.xy-kaka.cn'>基金推荐(alpha)</a>"
	if msg == "blog" || msg == "博客" {
		return personalBlog
	}
	if msg == "tech" || msg == "技术" || msg == "技术博客" {
		return technologyBlog
	}
	if msg == "基金" || msg == "fund" {
		return fundPage
	}
	
	return "功能列表: \n 1、" + personalBlog +
		" \n -----------------\n" +
		" 2、" +technologyBlog +
		" \n -----------------\n" +
		" 3、" + fundPage
}
