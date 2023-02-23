package tadmor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿沙拉AsharahBarony struct {
	feud.BaseBarony
}

var BAsharah阿沙拉 feud.Barony = &阿沙拉AsharahBarony{}

func init() {
	f := BAsharah阿沙拉.(*阿沙拉AsharahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asharah",
		TitleName: "阿沙拉",
		TitleCode: "b_asharah",
	}
}
