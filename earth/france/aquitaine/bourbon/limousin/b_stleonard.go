package limousin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣莱昂纳尔StleonardBarony struct {
	feud.BaseBarony
}

var BStleonard圣莱昂纳尔 feud.Barony = &圣莱昂纳尔StleonardBarony{}

func init() {
	f := BStleonard圣莱昂纳尔.(*圣莱昂纳尔StleonardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stleonard",
		TitleName: "圣莱昂纳尔",
		TitleCode: "b_stleonard",
	}
}
