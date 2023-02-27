package soyma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 季曼岭Timan_ridgeBarony struct {
	feud.BaseBarony
}

var BTiman_ridge季曼岭 feud.Barony = &季曼岭Timan_ridgeBarony{}

func init() {
    f := BTiman_ridge季曼岭.(*季曼岭Timan_ridgeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "timan_ridge",
		TitleName: "季曼岭",
		TitleCode: "b_timan_ridge",
	}
}
