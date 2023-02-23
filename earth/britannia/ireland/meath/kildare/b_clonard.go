package kildare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克洛纳德ClonardBarony struct {
	feud.BaseBarony
}

var BClonard克洛纳德 feud.Barony = &克洛纳德ClonardBarony{}

func init() {
	f := BClonard克洛纳德.(*克洛纳德ClonardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clonard",
		TitleName: "克洛纳德",
		TitleCode: "b_clonard",
	}
}
