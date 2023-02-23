package kalyani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨西尔SarseirBarony struct {
	feud.BaseBarony
}

var BSarseir萨西尔 feud.Barony = &萨西尔SarseirBarony{}

func init() {
	f := BSarseir萨西尔.(*萨西尔SarseirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarseir",
		TitleName: "萨西尔",
		TitleCode: "b_sarseir",
	}
}
