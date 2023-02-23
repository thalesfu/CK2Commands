package fars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卑路扎巴德PerozabadBarony struct {
	feud.BaseBarony
}

var BPerozabad卑路扎巴德 feud.Barony = &卑路扎巴德PerozabadBarony{}

func init() {
	f := BPerozabad卑路扎巴德.(*卑路扎巴德PerozabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perozabad",
		TitleName: "卑路扎巴德",
		TitleCode: "b_perozabad",
	}
}
