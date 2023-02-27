package macon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克吕尼ClunyBarony struct {
	feud.BaseBarony
}

var BCluny克吕尼 feud.Barony = &克吕尼ClunyBarony{}

func init() {
    f := BCluny克吕尼.(*克吕尼ClunyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cluny",
		TitleName: "克吕尼",
		TitleCode: "b_cluny",
	}
}
