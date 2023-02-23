package araouane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉万AraouaneBarony struct {
	feud.BaseBarony
}

var BAraouane阿拉万 feud.Barony = &阿拉万AraouaneBarony{}

func init() {
	f := BAraouane阿拉万.(*阿拉万AraouaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "araouane",
		TitleName: "阿拉万",
		TitleCode: "b_araouane",
	}
}
