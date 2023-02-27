package gobir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多巴DobaBarony struct {
	feud.BaseBarony
}

var BDoba多巴 feud.Barony = &多巴DobaBarony{}

func init() {
    f := BDoba多巴.(*多巴DobaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doba",
		TitleName: "多巴",
		TitleCode: "b_doba",
	}
}
