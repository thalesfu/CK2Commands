package smaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希尔SkirBarony struct {
	feud.BaseBarony
}

var BSkir希尔 feud.Barony = &希尔SkirBarony{}

func init() {
	f := BSkir希尔.(*希尔SkirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skir",
		TitleName: "希尔",
		TitleCode: "b_skir",
	}
}
