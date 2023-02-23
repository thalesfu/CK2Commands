package sasaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔胡ArhuBarony struct {
	feud.BaseBarony
}

var BArhu阿尔胡 feud.Barony = &阿尔胡ArhuBarony{}

func init() {
	f := BArhu阿尔胡.(*阿尔胡ArhuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arhu",
		TitleName: "阿尔胡",
		TitleCode: "b_arhu",
	}
}
