package farafra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因德拉Ain_dellaBarony struct {
	feud.BaseBarony
}

var BAin_della艾因德拉 feud.Barony = &艾因德拉Ain_dellaBarony{}

func init() {
    f := BAin_della艾因德拉.(*艾因德拉Ain_dellaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ain_della",
		TitleName: "艾因德拉",
		TitleCode: "b_ain_della",
	}
}
