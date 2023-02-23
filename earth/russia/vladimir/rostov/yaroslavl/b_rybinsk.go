package yaroslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷宾斯克RybinskBarony struct {
	feud.BaseBarony
}

var BRybinsk雷宾斯克 feud.Barony = &雷宾斯克RybinskBarony{}

func init() {
	f := BRybinsk雷宾斯克.(*雷宾斯克RybinskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rybinsk",
		TitleName: "雷宾斯克",
		TitleCode: "b_rybinsk",
	}
}
