package orania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥兰OranBarony struct {
	feud.BaseBarony
}

var BOran奥兰 feud.Barony = &奥兰OranBarony{}

func init() {
    f := BOran奥兰.(*奥兰OranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oran",
		TitleName: "奥兰",
		TitleCode: "b_oran",
	}
}
