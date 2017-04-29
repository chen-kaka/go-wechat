package processor

func ProcessText(msg string) string {
	if msg == "blog" || msg == "博客" {
		return "https://chen-kaka.github.io"
	}
	if msg == "基金" || msg == "fund" {
		return "http://fund.xy-kaka.cn"
	}
	
	return "谢谢关注。\n 功能列表:\n 博客 / blog \n 基金 / fund \n"
}
