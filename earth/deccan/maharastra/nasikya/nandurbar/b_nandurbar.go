package nandurbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 难豆罗婆罗NandurbarBarony struct {
	feud.BaseBarony
}

var BNandurbar难豆罗婆罗 feud.Barony = &难豆罗婆罗NandurbarBarony{}

func init() {
    f := BNandurbar难豆罗婆罗.(*难豆罗婆罗NandurbarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nandurbar",
		TitleName: "难豆罗婆罗",
		TitleCode: "b_nandurbar",
	}
}
