package shetland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶尔YellBarony struct {
	feud.BaseBarony
}

var BYell耶尔 feud.Barony = &耶尔YellBarony{}

func init() {
	f := BYell耶尔.(*耶尔YellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yell",
		TitleName: "耶尔",
		TitleCode: "b_yell",
	}
}
