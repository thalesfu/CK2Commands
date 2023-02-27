package hradec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日莱比ZlebyBarony struct {
	feud.BaseBarony
}

var BZleby日莱比 feud.Barony = &日莱比ZlebyBarony{}

func init() {
    f := BZleby日莱比.(*日莱比ZlebyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zleby",
		TitleName: "日莱比",
		TitleCode: "b_zleby",
	}
}
