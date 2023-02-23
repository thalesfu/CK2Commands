package frankfurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱茵费尔斯RheinfelsBarony struct {
	feud.BaseBarony
}

var BRheinfels莱茵费尔斯 feud.Barony = &莱茵费尔斯RheinfelsBarony{}

func init() {
	f := BRheinfels莱茵费尔斯.(*莱茵费尔斯RheinfelsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rheinfels",
		TitleName: "莱茵费尔斯",
		TitleCode: "b_rheinfels",
	}
}
