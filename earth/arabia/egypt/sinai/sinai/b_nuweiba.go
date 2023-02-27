package sinai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努韦巴NuweibaBarony struct {
	feud.BaseBarony
}

var BNuweiba努韦巴 feud.Barony = &努韦巴NuweibaBarony{}

func init() {
    f := BNuweiba努韦巴.(*努韦巴NuweibaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nuweiba",
		TitleName: "努韦巴",
		TitleCode: "b_nuweiba",
	}
}
