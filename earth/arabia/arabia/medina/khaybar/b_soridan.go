package khaybar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏里丹SoridanBarony struct {
	feud.BaseBarony
}

var BSoridan苏里丹 feud.Barony = &苏里丹SoridanBarony{}

func init() {
	f := BSoridan苏里丹.(*苏里丹SoridanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soridan",
		TitleName: "苏里丹",
		TitleCode: "b_soridan",
	}
}
