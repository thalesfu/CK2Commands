package manupura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坦泰TantaBarony struct {
	feud.BaseBarony
}

var BTanta坦泰 feud.Barony = &坦泰TantaBarony{}

func init() {
    f := BTanta坦泰.(*坦泰TantaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tanta",
		TitleName: "坦泰",
		TitleCode: "b_tanta",
	}
}
