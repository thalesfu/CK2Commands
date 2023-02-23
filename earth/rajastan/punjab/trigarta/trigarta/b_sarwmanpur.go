package trigarta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨婆摩那补罗SarwmanpurBarony struct {
	feud.BaseBarony
}

var BSarwmanpur萨婆摩那补罗 feud.Barony = &萨婆摩那补罗SarwmanpurBarony{}

func init() {
	f := BSarwmanpur萨婆摩那补罗.(*萨婆摩那补罗SarwmanpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarwmanpur",
		TitleName: "萨婆摩那补罗",
		TitleCode: "b_sarwmanpur",
	}
}
