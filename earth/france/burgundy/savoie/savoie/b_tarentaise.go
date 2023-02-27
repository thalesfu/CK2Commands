package savoie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔朗泰斯TarentaiseBarony struct {
	feud.BaseBarony
}

var BTarentaise塔朗泰斯 feud.Barony = &塔朗泰斯TarentaiseBarony{}

func init() {
    f := BTarentaise塔朗泰斯.(*塔朗泰斯TarentaiseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarentaise",
		TitleName: "塔朗泰斯",
		TitleCode: "b_tarentaise",
	}
}
