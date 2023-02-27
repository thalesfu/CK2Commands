package tsakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茶卡TsakhaBarony struct {
	feud.BaseBarony
}

var BTsakha茶卡 feud.Barony = &茶卡TsakhaBarony{}

func init() {
    f := BTsakha茶卡.(*茶卡TsakhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsakha",
		TitleName: "茶卡",
		TitleCode: "b_tsakha",
	}
}
