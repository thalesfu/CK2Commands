package farrah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉达鲁克BaladulukBarony struct {
	feud.BaseBarony
}

var BBaladuluk巴拉达鲁克 feud.Barony = &巴拉达鲁克BaladulukBarony{}

func init() {
    f := BBaladuluk巴拉达鲁克.(*巴拉达鲁克BaladulukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baladuluk",
		TitleName: "巴拉达鲁克",
		TitleCode: "b_baladuluk",
	}
}
