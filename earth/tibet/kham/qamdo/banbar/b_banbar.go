package banbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 边坝BanbarBarony struct {
	feud.BaseBarony
}

var BBanbar边坝 feud.Barony = &边坝BanbarBarony{}

func init() {
    f := BBanbar边坝.(*边坝BanbarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banbar",
		TitleName: "边坝",
		TitleCode: "b_banbar",
	}
}
