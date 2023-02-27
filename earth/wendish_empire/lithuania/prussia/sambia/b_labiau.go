package sambia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉比奥LabiauBarony struct {
	feud.BaseBarony
}

var BLabiau拉比奥 feud.Barony = &拉比奥LabiauBarony{}

func init() {
    f := BLabiau拉比奥.(*拉比奥LabiauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "labiau",
		TitleName: "拉比奥",
		TitleCode: "b_labiau",
	}
}
