package zachlumia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉奇BracBarony struct {
	feud.BaseBarony
}

var BBrac布拉奇 feud.Barony = &布拉奇BracBarony{}

func init() {
    f := BBrac布拉奇.(*布拉奇BracBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brac",
		TitleName: "布拉奇",
		TitleCode: "b_brac",
	}
}
