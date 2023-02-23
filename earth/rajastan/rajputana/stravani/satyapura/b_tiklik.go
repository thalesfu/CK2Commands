package satyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 置力TiklikBarony struct {
	feud.BaseBarony
}

var BTiklik置力 feud.Barony = &置力TiklikBarony{}

func init() {
	f := BTiklik置力.(*置力TiklikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiklik",
		TitleName: "置力",
		TitleCode: "b_tiklik",
	}
}
