package vatsagulma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩奴罗ManoraBarony struct {
	feud.BaseBarony
}

var BManora摩奴罗 feud.Barony = &摩奴罗ManoraBarony{}

func init() {
	f := BManora摩奴罗.(*摩奴罗ManoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manora",
		TitleName: "摩奴罗",
		TitleCode: "b_manora",
	}
}
