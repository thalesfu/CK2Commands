package khozistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴丹AbadanBarony struct {
	feud.BaseBarony
}

var BAbadan阿巴丹 feud.Barony = &阿巴丹AbadanBarony{}

func init() {
	f := BAbadan阿巴丹.(*阿巴丹AbadanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abadan",
		TitleName: "阿巴丹",
		TitleCode: "b_abadan",
	}
}
