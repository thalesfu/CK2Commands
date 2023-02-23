package nyssa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿耳刻莱斯ArchelaisBarony struct {
	feud.BaseBarony
}

var BArchelais阿耳刻莱斯 feud.Barony = &阿耳刻莱斯ArchelaisBarony{}

func init() {
	f := BArchelais阿耳刻莱斯.(*阿耳刻莱斯ArchelaisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "archelais",
		TitleName: "阿耳刻莱斯",
		TitleCode: "b_archelais",
	}
}
