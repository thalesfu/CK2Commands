package burtasy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈雷卡拉梅什GolykaramyshBarony struct {
	feud.BaseBarony
}

var BGolykaramysh戈雷卡拉梅什 feud.Barony = &戈雷卡拉梅什GolykaramyshBarony{}

func init() {
    f := BGolykaramysh戈雷卡拉梅什.(*戈雷卡拉梅什GolykaramyshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "golykaramysh",
		TitleName: "戈雷卡拉梅什",
		TitleCode: "b_golykaramysh",
	}
}
