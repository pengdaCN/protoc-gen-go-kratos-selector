package main

import (
	"bytes"
	"testing"
)

func TestExec(t *testing.T) {
	var bs = bytes.NewBuffer(nil)

	data := DrawingData{
		Name: "test1",
		//Middleware: kratosMiddleware,
		//Selector:   kratosSelector,
		Handlers: []struct {
			Name           string
			FullMethodName []string
		}{
			{
				Name: "jwt",
				FullMethodName: []string{
					"/fdsa/",
					"/.fdsa",
				},
			},
		},
	}

	err := tmpl.Execute(bs, &data)
	if err != nil {
		t.Fatal(err)
	}

	//t.Log(kratosSelector.Ident("Selector").String())

	t.Log(bs.String())
}
