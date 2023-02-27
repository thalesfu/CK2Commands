package taghaza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔塔AlouartaBarony struct {
	feud.BaseBarony
}

var BAlouarta瓦尔塔 feud.Barony = &瓦尔塔AlouartaBarony{}

func init() {
    f := BAlouarta瓦尔塔.(*瓦尔塔AlouartaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alouarta",
		TitleName: "瓦尔塔",
		TitleCode: "b_alouarta",
	}
}
