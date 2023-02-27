package hedmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旺VangBarony struct {
	feud.BaseBarony
}

var BVang旺 feud.Barony = &旺VangBarony{}

func init() {
    f := BVang旺.(*旺VangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vang",
		TitleName: "旺",
		TitleCode: "b_vang",
	}
}
