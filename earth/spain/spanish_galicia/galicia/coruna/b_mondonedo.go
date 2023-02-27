package coruna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙多涅多MondonedoBarony struct {
	feud.BaseBarony
}

var BMondonedo蒙多涅多 feud.Barony = &蒙多涅多MondonedoBarony{}

func init() {
    f := BMondonedo蒙多涅多.(*蒙多涅多MondonedoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mondonedo",
		TitleName: "蒙多涅多",
		TitleCode: "b_mondonedo",
	}
}
