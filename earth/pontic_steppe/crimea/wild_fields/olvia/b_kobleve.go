package olvia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科布列韦KobleveBarony struct {
	feud.BaseBarony
}

var BKobleve科布列韦 feud.Barony = &科布列韦KobleveBarony{}

func init() {
    f := BKobleve科布列韦.(*科布列韦KobleveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kobleve",
		TitleName: "科布列韦",
		TitleCode: "b_kobleve",
	}
}
