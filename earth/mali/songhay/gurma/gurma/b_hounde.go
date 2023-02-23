package gurma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洪代HoundeBarony struct {
	feud.BaseBarony
}

var BHounde洪代 feud.Barony = &洪代HoundeBarony{}

func init() {
	f := BHounde洪代.(*洪代HoundeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hounde",
		TitleName: "洪代",
		TitleCode: "b_hounde",
	}
}
