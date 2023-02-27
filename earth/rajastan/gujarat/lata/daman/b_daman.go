package daman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀曼那DamanBarony struct {
	feud.BaseBarony
}

var BDaman陀曼那 feud.Barony = &陀曼那DamanBarony{}

func init() {
    f := BDaman陀曼那.(*陀曼那DamanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daman",
		TitleName: "陀曼那",
		TitleCode: "b_daman",
	}
}
