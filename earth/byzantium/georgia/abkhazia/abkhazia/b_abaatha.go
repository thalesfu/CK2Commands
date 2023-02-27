package abkhazia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴特哈AbaathaBarony struct {
	feud.BaseBarony
}

var BAbaatha阿巴特哈 feud.Barony = &阿巴特哈AbaathaBarony{}

func init() {
    f := BAbaatha阿巴特哈.(*阿巴特哈AbaathaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abaatha",
		TitleName: "阿巴特哈",
		TitleCode: "b_abaatha",
	}
}
