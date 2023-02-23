package klarjeti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉尔哲季KlarjetiBarony struct {
	feud.BaseBarony
}

var BKlarjeti克拉尔哲季 feud.Barony = &克拉尔哲季KlarjetiBarony{}

func init() {
	f := BKlarjeti克拉尔哲季.(*克拉尔哲季KlarjetiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klarjeti",
		TitleName: "克拉尔哲季",
		TitleCode: "b_klarjeti",
	}
}
