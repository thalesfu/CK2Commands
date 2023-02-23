package strymon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉多维什RadovishBarony struct {
	feud.BaseBarony
}

var BRadovish拉多维什 feud.Barony = &拉多维什RadovishBarony{}

func init() {
	f := BRadovish拉多维什.(*拉多维什RadovishBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radovish",
		TitleName: "拉多维什",
		TitleCode: "b_radovish",
	}
}
