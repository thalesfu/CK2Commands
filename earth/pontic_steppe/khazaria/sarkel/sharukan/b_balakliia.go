package sharukan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉克列亚BalakliiaBarony struct {
	feud.BaseBarony
}

var BBalakliia巴拉克列亚 feud.Barony = &巴拉克列亚BalakliiaBarony{}

func init() {
    f := BBalakliia巴拉克列亚.(*巴拉克列亚BalakliiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balakliia",
		TitleName: "巴拉克列亚",
		TitleCode: "b_balakliia",
	}
}
