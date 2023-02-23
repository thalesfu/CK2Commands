package gar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甲岗ChakgangBarony struct {
	feud.BaseBarony
}

var BChakgang甲岗 feud.Barony = &甲岗ChakgangBarony{}

func init() {
	f := BChakgang甲岗.(*甲岗ChakgangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chakgang",
		TitleName: "甲岗",
		TitleCode: "b_chakgang",
	}
}
