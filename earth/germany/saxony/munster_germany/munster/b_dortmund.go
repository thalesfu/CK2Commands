package munster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多特蒙德DortmundBarony struct {
	feud.BaseBarony
}

var BDortmund多特蒙德 feud.Barony = &多特蒙德DortmundBarony{}

func init() {
    f := BDortmund多特蒙德.(*多特蒙德DortmundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dortmund",
		TitleName: "多特蒙德",
		TitleCode: "b_dortmund",
	}
}
