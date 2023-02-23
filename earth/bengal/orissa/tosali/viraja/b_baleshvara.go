package viraja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆丽湿伐罗BaleshvaraBarony struct {
	feud.BaseBarony
}

var BBaleshvara婆丽湿伐罗 feud.Barony = &婆丽湿伐罗BaleshvaraBarony{}

func init() {
	f := BBaleshvara婆丽湿伐罗.(*婆丽湿伐罗BaleshvaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baleshvara",
		TitleName: "婆丽湿伐罗",
		TitleCode: "b_baleshvara",
	}
}
