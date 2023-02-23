package lleida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉格尔BalaguerBarony struct {
	feud.BaseBarony
}

var BBalaguer巴拉格尔 feud.Barony = &巴拉格尔BalaguerBarony{}

func init() {
	f := BBalaguer巴拉格尔.(*巴拉格尔BalaguerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balaguer",
		TitleName: "巴拉格尔",
		TitleCode: "b_balaguer",
	}
}
