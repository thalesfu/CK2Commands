package desht_i_kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴赫穆特BakhmutBarony struct {
	feud.BaseBarony
}

var BBakhmut巴赫穆特 feud.Barony = &巴赫穆特BakhmutBarony{}

func init() {
    f := BBakhmut巴赫穆特.(*巴赫穆特BakhmutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakhmut",
		TitleName: "巴赫穆特",
		TitleCode: "b_bakhmut",
	}
}
