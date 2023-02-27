package kiev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦西里基夫VasyklivBarony struct {
	feud.BaseBarony
}

var BVasykliv瓦西里基夫 feud.Barony = &瓦西里基夫VasyklivBarony{}

func init() {
    f := BVasykliv瓦西里基夫.(*瓦西里基夫VasyklivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vasykliv",
		TitleName: "瓦西里基夫",
		TitleCode: "b_vasykliv",
	}
}
