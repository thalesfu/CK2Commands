package makran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔马拉OrmaraBarony struct {
	feud.BaseBarony
}

var BOrmara奥尔马拉 feud.Barony = &奥尔马拉OrmaraBarony{}

func init() {
    f := BOrmara奥尔马拉.(*奥尔马拉OrmaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ormara",
		TitleName: "奥尔马拉",
		TitleCode: "b_ormara",
	}
}
