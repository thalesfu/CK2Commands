package otrar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朔什卡凯ShoshkakoiBarony struct {
	feud.BaseBarony
}

var BShoshkakoi朔什卡凯 feud.Barony = &朔什卡凯ShoshkakoiBarony{}

func init() {
    f := BShoshkakoi朔什卡凯.(*朔什卡凯ShoshkakoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shoshkakoi",
		TitleName: "朔什卡凯",
		TitleCode: "b_shoshkakoi",
	}
}
