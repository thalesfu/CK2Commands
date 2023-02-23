package eu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福尔日ForgesBarony struct {
	feud.BaseBarony
}

var BForges福尔日 feud.Barony = &福尔日ForgesBarony{}

func init() {
	f := BForges福尔日.(*福尔日ForgesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "forges",
		TitleName: "福尔日",
		TitleCode: "b_forges",
	}
}
