package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("%s,%s", p.Name, p.Name)
}
func main() {
	p := &People{Name: "123"}
	fmt.Println(p.String())
}
