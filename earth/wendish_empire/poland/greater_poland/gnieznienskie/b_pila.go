package gnieznienskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮瓦PilaBarony struct {
	feud.BaseBarony
}

var BPila皮瓦 feud.Barony = &皮瓦PilaBarony{}

func init() {
    f := BPila皮瓦.(*皮瓦PilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pila",
		TitleName: "皮瓦",
		TitleCode: "b_pila",
	}
}
