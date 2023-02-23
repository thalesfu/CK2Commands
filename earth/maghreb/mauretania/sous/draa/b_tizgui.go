package draa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提兹古TizguiBarony struct {
	feud.BaseBarony
}

var BTizgui提兹古 feud.Barony = &提兹古TizguiBarony{}

func init() {
	f := BTizgui提兹古.(*提兹古TizguiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tizgui",
		TitleName: "提兹古",
		TitleCode: "b_tizgui",
	}
}
