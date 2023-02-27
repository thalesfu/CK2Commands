package tyrconnell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴利马克斯万尼BallymacswineyBarony struct {
	feud.BaseBarony
}

var BBallymacswiney巴利马克斯万尼 feud.Barony = &巴利马克斯万尼BallymacswineyBarony{}

func init() {
    f := BBallymacswiney巴利马克斯万尼.(*巴利马克斯万尼BallymacswineyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ballymacswiney",
		TitleName: "巴利马克斯万尼",
		TitleCode: "b_ballymacswiney",
	}
}
