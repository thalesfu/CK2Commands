package minden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍克莱韦HockeleveBarony struct {
	feud.BaseBarony
}

var BHockeleve霍克莱韦 feud.Barony = &霍克莱韦HockeleveBarony{}

func init() {
    f := BHockeleve霍克莱韦.(*霍克莱韦HockeleveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hockeleve",
		TitleName: "霍克莱韦",
		TitleCode: "b_hockeleve",
	}
}
