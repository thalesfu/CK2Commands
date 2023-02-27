package velikiye_luki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢别日SebezhBarony struct {
	feud.BaseBarony
}

var BSebezh谢别日 feud.Barony = &谢别日SebezhBarony{}

func init() {
    f := BSebezh谢别日.(*谢别日SebezhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sebezh",
		TitleName: "谢别日",
		TitleCode: "b_sebezh",
	}
}
