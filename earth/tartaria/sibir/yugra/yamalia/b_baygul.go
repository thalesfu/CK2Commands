package yamalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜古尔BaygulBarony struct {
	feud.BaseBarony
}

var BBaygul拜古尔 feud.Barony = &拜古尔BaygulBarony{}

func init() {
    f := BBaygul拜古尔.(*拜古尔BaygulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baygul",
		TitleName: "拜古尔",
		TitleCode: "b_baygul",
	}
}
