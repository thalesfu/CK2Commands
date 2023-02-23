package wadai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于埃YueBarony struct {
	feud.BaseBarony
}

var BYue于埃 feud.Barony = &于埃YueBarony{}

func init() {
	f := BYue于埃.(*于埃YueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yue",
		TitleName: "于埃",
		TitleCode: "b_yue",
	}
}
