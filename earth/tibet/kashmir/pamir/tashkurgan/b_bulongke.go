package tashkurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布隆克BulongkeBarony struct {
	feud.BaseBarony
}

var BBulongke布隆克 feud.Barony = &布隆克BulongkeBarony{}

func init() {
    f := BBulongke布隆克.(*布隆克BulongkeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bulongke",
		TitleName: "布隆克",
		TitleCode: "b_bulongke",
	}
}
