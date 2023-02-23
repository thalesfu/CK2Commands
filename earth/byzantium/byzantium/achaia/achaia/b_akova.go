package achaia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿科瓦AkovaBarony struct {
	feud.BaseBarony
}

var BAkova阿科瓦 feud.Barony = &阿科瓦AkovaBarony{}

func init() {
	f := BAkova阿科瓦.(*阿科瓦AkovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akova",
		TitleName: "阿科瓦",
		TitleCode: "b_akova",
	}
}
