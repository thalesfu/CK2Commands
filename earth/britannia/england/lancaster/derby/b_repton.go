package derby

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷普顿ReptonBarony struct {
	feud.BaseBarony
}

var BRepton雷普顿 feud.Barony = &雷普顿ReptonBarony{}

func init() {
	f := BRepton雷普顿.(*雷普顿ReptonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "repton",
		TitleName: "雷普顿",
		TitleCode: "b_repton",
	}
}
