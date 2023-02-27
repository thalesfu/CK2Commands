package ghazna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲈鳗LomanBarony struct {
	feud.BaseBarony
}

var BLoman鲈鳗 feud.Barony = &鲈鳗LomanBarony{}

func init() {
    f := BLoman鲈鳗.(*鲈鳗LomanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loman",
		TitleName: "鲈鳗",
		TitleCode: "b_loman",
	}
}
