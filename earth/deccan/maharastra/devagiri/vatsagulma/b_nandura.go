package vatsagulma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 难豆罗NanduraBarony struct {
	feud.BaseBarony
}

var BNandura难豆罗 feud.Barony = &难豆罗NanduraBarony{}

func init() {
	f := BNandura难豆罗.(*难豆罗NanduraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nandura",
		TitleName: "难豆罗",
		TitleCode: "b_nandura",
	}
}
