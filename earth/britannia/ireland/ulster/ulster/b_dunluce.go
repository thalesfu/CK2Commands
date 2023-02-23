package ulster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓卢斯DunluceBarony struct {
	feud.BaseBarony
}

var BDunluce邓卢斯 feud.Barony = &邓卢斯DunluceBarony{}

func init() {
	f := BDunluce邓卢斯.(*邓卢斯DunluceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunluce",
		TitleName: "邓卢斯",
		TitleCode: "b_dunluce",
	}
}
