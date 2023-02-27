package zaranj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米拉克MilakBarony struct {
	feud.BaseBarony
}

var BMilak米拉克 feud.Barony = &米拉克MilakBarony{}

func init() {
    f := BMilak米拉克.(*米拉克MilakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "milak",
		TitleName: "米拉克",
		TitleCode: "b_milak",
	}
}
