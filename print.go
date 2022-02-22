package golor

import "fmt"

type Color string

const (
	BLACK   Color = "\033[1;30m%s\033[0m"
	RED     Color = "\033[1;31m%s\033[0m"
	GREEN   Color = "\033[1;32m%s\033[0m"
	YELLOW  Color = "\033[1;33m%s\033[0m"
	PURPLE  Color = "\033[1;34m%s\033[0m"
	MAGENTA Color = "\033[1;35m%s\033[0m"
	TEAL    Color = "\033[1;36m%s\033[0m"
	WHITE   Color = "\033[1;37m%s\033[0m"
)

func Printf(c Color, format string, a ...interface{}) {
	fmt.Printf(string(c), fmt.Sprintf(format, a...))
}
func Sprintf(c Color, format string, a ...interface{}) string {
	return fmt.Sprintf(string(c), fmt.Sprintf(format, a...))
}
func Red(format string) string {
	return fmt.Sprintf(string(RED), format)
}
func Black(format string) string {
	return fmt.Sprintf(string(BLACK), format)
}
