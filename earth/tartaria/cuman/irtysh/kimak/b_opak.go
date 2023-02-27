package kimak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥帕克OpakBarony struct {
	feud.BaseBarony
}

var BOpak奥帕克 feud.Barony = &奥帕克OpakBarony{}

func init() {
    f := BOpak奥帕克.(*奥帕克OpakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "opak",
		TitleName: "奥帕克",
		TitleCode: "b_opak",
	}
}
