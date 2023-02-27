package lyon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普西尼昂PusignanBarony struct {
	feud.BaseBarony
}

var BPusignan普西尼昂 feud.Barony = &普西尼昂PusignanBarony{}

func init() {
    f := BPusignan普西尼昂.(*普西尼昂PusignanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pusignan",
		TitleName: "普西尼昂",
		TitleCode: "b_pusignan",
	}
}
