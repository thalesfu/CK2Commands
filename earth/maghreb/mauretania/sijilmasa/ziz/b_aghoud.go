package ziz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿古德AghoudBarony struct {
	feud.BaseBarony
}

var BAghoud阿古德 feud.Barony = &阿古德AghoudBarony{}

func init() {
	f := BAghoud阿古德.(*阿古德AghoudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aghoud",
		TitleName: "阿古德",
		TitleCode: "b_aghoud",
	}
}
