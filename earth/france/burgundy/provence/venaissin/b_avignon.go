package venaissin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿维尼翁AvignonBarony struct {
	feud.BaseBarony
}

var BAvignon阿维尼翁 feud.Barony = &阿维尼翁AvignonBarony{}

func init() {
	f := BAvignon阿维尼翁.(*阿维尼翁AvignonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avignon",
		TitleName: "阿维尼翁",
		TitleCode: "b_avignon",
	}
}
