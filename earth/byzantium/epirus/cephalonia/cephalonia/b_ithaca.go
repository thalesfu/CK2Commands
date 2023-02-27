package cephalonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊萨卡IthacaBarony struct {
	feud.BaseBarony
}

var BIthaca伊萨卡 feud.Barony = &伊萨卡IthacaBarony{}

func init() {
    f := BIthaca伊萨卡.(*伊萨卡IthacaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ithaca",
		TitleName: "伊萨卡",
		TitleCode: "b_ithaca",
	}
}
