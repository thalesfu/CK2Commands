package theodosia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥里瓦OlyvaBarony struct {
	feud.BaseBarony
}

var BOlyva奥里瓦 feud.Barony = &奥里瓦OlyvaBarony{}

func init() {
    f := BOlyva奥里瓦.(*奥里瓦OlyvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olyva",
		TitleName: "奥里瓦",
		TitleCode: "b_olyva",
	}
}
