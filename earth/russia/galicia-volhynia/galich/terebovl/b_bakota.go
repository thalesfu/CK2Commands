package terebovl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴科塔BakotaBarony struct {
	feud.BaseBarony
}

var BBakota巴科塔 feud.Barony = &巴科塔BakotaBarony{}

func init() {
    f := BBakota巴科塔.(*巴科塔BakotaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakota",
		TitleName: "巴科塔",
		TitleCode: "b_bakota",
	}
}
