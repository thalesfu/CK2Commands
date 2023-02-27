package bulgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 保加尔BulgarBarony struct {
	feud.BaseBarony
}

var BBulgar保加尔 feud.Barony = &保加尔BulgarBarony{}

func init() {
    f := BBulgar保加尔.(*保加尔BulgarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bulgar",
		TitleName: "保加尔",
		TitleCode: "b_bulgar",
	}
}
