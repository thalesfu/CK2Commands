package bergenshus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欣萨维克KinsarvikBarony struct {
	feud.BaseBarony
}

var BKinsarvik欣萨维克 feud.Barony = &欣萨维克KinsarvikBarony{}

func init() {
	f := BKinsarvik欣萨维克.(*欣萨维克KinsarvikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kinsarvik",
		TitleName: "欣萨维克",
		TitleCode: "b_kinsarvik",
	}
}
