package brycheiniog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉多克CradocBarony struct {
	feud.BaseBarony
}

var BCradoc克拉多克 feud.Barony = &克拉多克CradocBarony{}

func init() {
    f := BCradoc克拉多克.(*克拉多克CradocBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cradoc",
		TitleName: "克拉多克",
		TitleCode: "b_cradoc",
	}
}
