package aracena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加拉罗萨GalarozaBarony struct {
	feud.BaseBarony
}

var BGalaroza加拉罗萨 feud.Barony = &加拉罗萨GalarozaBarony{}

func init() {
    f := BGalaroza加拉罗萨.(*加拉罗萨GalarozaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galaroza",
		TitleName: "加拉罗萨",
		TitleCode: "b_galaroza",
	}
}
