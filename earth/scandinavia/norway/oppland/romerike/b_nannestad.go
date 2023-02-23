package romerike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南内斯塔NannestadBarony struct {
	feud.BaseBarony
}

var BNannestad南内斯塔 feud.Barony = &南内斯塔NannestadBarony{}

func init() {
	f := BNannestad南内斯塔.(*南内斯塔NannestadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nannestad",
		TitleName: "南内斯塔",
		TitleCode: "b_nannestad",
	}
}
