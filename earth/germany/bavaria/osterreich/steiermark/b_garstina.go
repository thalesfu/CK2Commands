package steiermark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加斯蒂纳GarstinaBarony struct {
	feud.BaseBarony
}

var BGarstina加斯蒂纳 feud.Barony = &加斯蒂纳GarstinaBarony{}

func init() {
	f := BGarstina加斯蒂纳.(*加斯蒂纳GarstinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "garstina",
		TitleName: "加斯蒂纳",
		TitleCode: "b_garstina",
	}
}
