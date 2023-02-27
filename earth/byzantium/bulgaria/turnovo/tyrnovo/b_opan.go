package tyrnovo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥潘OpanBarony struct {
	feud.BaseBarony
}

var BOpan奥潘 feud.Barony = &奥潘OpanBarony{}

func init() {
    f := BOpan奥潘.(*奥潘OpanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "opan",
		TitleName: "奥潘",
		TitleCode: "b_opan",
	}
}
