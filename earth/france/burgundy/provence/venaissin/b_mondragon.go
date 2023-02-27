package venaissin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙德拉贡MondragonBarony struct {
	feud.BaseBarony
}

var BMondragon蒙德拉贡 feud.Barony = &蒙德拉贡MondragonBarony{}

func init() {
    f := BMondragon蒙德拉贡.(*蒙德拉贡MondragonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mondragon",
		TitleName: "蒙德拉贡",
		TitleCode: "b_mondragon",
	}
}
