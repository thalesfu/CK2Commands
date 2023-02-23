package cagliari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿塞米尼AsseminiBarony struct {
	feud.BaseBarony
}

var BAssemini阿塞米尼 feud.Barony = &阿塞米尼AsseminiBarony{}

func init() {
	f := BAssemini阿塞米尼.(*阿塞米尼AsseminiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "assemini",
		TitleName: "阿塞米尼",
		TitleCode: "b_assemini",
	}
}
