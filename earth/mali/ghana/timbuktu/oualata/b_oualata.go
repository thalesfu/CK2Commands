package oualata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦拉塔OualataBarony struct {
	feud.BaseBarony
}

var BOualata瓦拉塔 feud.Barony = &瓦拉塔OualataBarony{}

func init() {
    f := BOualata瓦拉塔.(*瓦拉塔OualataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oualata",
		TitleName: "瓦拉塔",
		TitleCode: "b_oualata",
	}
}
