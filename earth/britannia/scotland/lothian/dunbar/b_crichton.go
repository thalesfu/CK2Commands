package dunbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克赖顿CrichtonBarony struct {
	feud.BaseBarony
}

var BCrichton克赖顿 feud.Barony = &克赖顿CrichtonBarony{}

func init() {
    f := BCrichton克赖顿.(*克赖顿CrichtonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crichton",
		TitleName: "克赖顿",
		TitleCode: "b_crichton",
	}
}
