package rostov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯托夫RostovBarony struct {
	feud.BaseBarony
}

var BRostov罗斯托夫 feud.Barony = &罗斯托夫RostovBarony{}

func init() {
	f := BRostov罗斯托夫.(*罗斯托夫RostovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rostov",
		TitleName: "罗斯托夫",
		TitleCode: "b_rostov",
	}
}
