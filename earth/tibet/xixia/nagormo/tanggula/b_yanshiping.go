package tanggula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雁石坪YanshipingBarony struct {
	feud.BaseBarony
}

var BYanshiping雁石坪 feud.Barony = &雁石坪YanshipingBarony{}

func init() {
	f := BYanshiping雁石坪.(*雁石坪YanshipingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yanshiping",
		TitleName: "雁石坪",
		TitleCode: "b_yanshiping",
	}
}
