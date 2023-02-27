package farafra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因比射Ain_bishayBarony struct {
	feud.BaseBarony
}

var BAin_bishay艾因比射 feud.Barony = &艾因比射Ain_bishayBarony{}

func init() {
    f := BAin_bishay艾因比射.(*艾因比射Ain_bishayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ain_bishay",
		TitleName: "艾因比射",
		TitleCode: "b_ain_bishay",
	}
}
