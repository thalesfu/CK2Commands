package chios

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊卡利亚IkariaBarony struct {
	feud.BaseBarony
}

var BIkaria伊卡利亚 feud.Barony = &伊卡利亚IkariaBarony{}

func init() {
	f := BIkaria伊卡利亚.(*伊卡利亚IkariaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ikaria",
		TitleName: "伊卡利亚",
		TitleCode: "b_ikaria",
	}
}
