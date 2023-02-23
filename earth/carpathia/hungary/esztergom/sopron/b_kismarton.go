package sopron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基什毛尔通KismartonBarony struct {
	feud.BaseBarony
}

var BKismarton基什毛尔通 feud.Barony = &基什毛尔通KismartonBarony{}

func init() {
	f := BKismarton基什毛尔通.(*基什毛尔通KismartonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kismarton",
		TitleName: "基什毛尔通",
		TitleCode: "b_kismarton",
	}
}
