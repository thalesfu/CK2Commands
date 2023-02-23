package nandapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 难陀城NandapurBarony struct {
	feud.BaseBarony
}

var BNandapur难陀城 feud.Barony = &难陀城NandapurBarony{}

func init() {
	f := BNandapur难陀城.(*难陀城NandapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nandapur",
		TitleName: "难陀城",
		TitleCode: "b_nandapur",
	}
}
