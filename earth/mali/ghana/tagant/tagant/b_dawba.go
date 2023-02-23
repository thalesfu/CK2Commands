package tagant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 道巴DawbaBarony struct {
	feud.BaseBarony
}

var BDawba道巴 feud.Barony = &道巴DawbaBarony{}

func init() {
	f := BDawba道巴.(*道巴DawbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dawba",
		TitleName: "道巴",
		TitleCode: "b_dawba",
	}
}
