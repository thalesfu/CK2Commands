package irgiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔KairBarony struct {
	feud.BaseBarony
}

var BKair凯尔 feud.Barony = &凯尔KairBarony{}

func init() {
	f := BKair凯尔.(*凯尔KairBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kair",
		TitleName: "凯尔",
		TitleCode: "b_kair",
	}
}
