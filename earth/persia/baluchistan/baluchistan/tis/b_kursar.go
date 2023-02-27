package tis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库雷斯萨KursarBarony struct {
	feud.BaseBarony
}

var BKursar库雷斯萨 feud.Barony = &库雷斯萨KursarBarony{}

func init() {
    f := BKursar库雷斯萨.(*库雷斯萨KursarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kursar",
		TitleName: "库雷斯萨",
		TitleCode: "b_kursar",
	}
}
