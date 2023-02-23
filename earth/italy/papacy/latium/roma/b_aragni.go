package roma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿纳尼AragniBarony struct {
	feud.BaseBarony
}

var BAragni阿纳尼 feud.Barony = &阿纳尼AragniBarony{}

func init() {
	f := BAragni阿纳尼.(*阿纳尼AragniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aragni",
		TitleName: "阿纳尼",
		TitleCode: "b_aragni",
	}
}
