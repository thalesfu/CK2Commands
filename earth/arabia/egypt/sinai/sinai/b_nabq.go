package sinai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈卜格NabqBarony struct {
	feud.BaseBarony
}

var BNabq奈卜格 feud.Barony = &奈卜格NabqBarony{}

func init() {
	f := BNabq奈卜格.(*奈卜格NabqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nabq",
		TitleName: "奈卜格",
		TitleCode: "b_nabq",
	}
}
