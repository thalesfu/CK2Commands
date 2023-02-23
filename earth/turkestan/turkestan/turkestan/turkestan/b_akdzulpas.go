package turkestan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克_祖尔帕斯AkdzulpasBarony struct {
	feud.BaseBarony
}

var BAkdzulpas阿克_祖尔帕斯 feud.Barony = &阿克_祖尔帕斯AkdzulpasBarony{}

func init() {
	f := BAkdzulpas阿克_祖尔帕斯.(*阿克_祖尔帕斯AkdzulpasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akdzulpas",
		TitleName: "阿克_祖尔帕斯",
		TitleCode: "b_akdzulpas",
	}
}
