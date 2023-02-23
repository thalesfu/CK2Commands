package om

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯托夫卡RostovkaBarony struct {
	feud.BaseBarony
}

var BRostovka罗斯托夫卡 feud.Barony = &罗斯托夫卡RostovkaBarony{}

func init() {
	f := BRostovka罗斯托夫卡.(*罗斯托夫卡RostovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rostovka",
		TitleName: "罗斯托夫卡",
		TitleCode: "b_rostovka",
	}
}
