package abakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴坎AbakanBarony struct {
	feud.BaseBarony
}

var BAbakan阿巴坎 feud.Barony = &阿巴坎AbakanBarony{}

func init() {
	f := BAbakan阿巴坎.(*阿巴坎AbakanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abakan",
		TitleName: "阿巴坎",
		TitleCode: "b_abakan",
	}
}
