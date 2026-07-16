package utils

import (
	"fmt"
	"testing"
)

type S struct {
	N string
}

type F struct {
	UserName string
	S        S
}

func TestHTML(t *testing.T) {
	a := `
<!DOCTYPE html>
<html>
<head>
	<title>template</title>
</head>
<body>
hello {{.UserName}}<br>
hello {{.S.N}}<br>
</body>
</html>
	`
	f := F{UserName: "aaaa", S: S{
		N: "nnnn",
	}}
	fmt.Println(HTML(a, f))
}
