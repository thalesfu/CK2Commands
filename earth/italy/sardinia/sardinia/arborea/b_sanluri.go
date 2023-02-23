package arborea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣卢里SanluriBarony struct {
	feud.BaseBarony
}

var BSanluri圣卢里 feud.Barony = &圣卢里SanluriBarony{}

func init() {
	f := BSanluri圣卢里.(*圣卢里SanluriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanluri",
		TitleName: "圣卢里",
		TitleCode: "b_sanluri",
	}
}
