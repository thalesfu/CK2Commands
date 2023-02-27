package al_aqabah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古韦拉QuwairahBarony struct {
	feud.BaseBarony
}

var BQuwairah古韦拉 feud.Barony = &古韦拉QuwairahBarony{}

func init() {
    f := BQuwairah古韦拉.(*古韦拉QuwairahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quwairah",
		TitleName: "古韦拉",
		TitleCode: "b_quwairah",
	}
}
