package basics

import (
	"fmt"
	"os"
	"testing"
	"text/template"
)

type Person1 struct {
	Name                string
	nonExportedAgeField string
}

func TestTemplateDemo(tt *testing.T) {
	tmp := template.New("hello")
	t, e := tmp.Parse("hello {{.Name}}!\n")
	// 校验模板格式，不正确会panic
	template.Must(t, e)
	p := Person1{Name: "Mary", nonExportedAgeField: "31"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	//if else
	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} Print IF part. {{else}} Print ELSE part.{{end}}\n")) //non empty pipeline following if condition
	tIfElse.Execute(os.Stdout, nil)

	//with-end ：会把with指定的值推送给点号{{.}}，点号就是个占位符
	t, _ = t.Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}}{{end}}!\n")
	t.Execute(os.Stdout, nil)

	//变量
	t = template.Must(t.Parse("{{with $x_1 := `hey`}}{{$x_1}} {{.}} {{$x_1}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)

	//range-end
	s := []int{1, 2, 3, 4}
	t = template.Must(t.Parse("{{range .}} {{.}} {{else}} arr is null {{end}}\n"))
	t.Execute(os.Stdout, s)
	t.Execute(os.Stdout, nil)

	// printf
	t = template.Must(t.Parse("{{with $x := `hello`}}{{printf `%s %s` $x `Mary`}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)
}
