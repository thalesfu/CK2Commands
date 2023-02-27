package la_marche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝拉克BellacBarony struct {
	feud.BaseBarony
}

var BBellac贝拉克 feud.Barony = &贝拉克BellacBarony{}

func init() {
    f := BBellac贝拉克.(*贝拉克BellacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bellac",
		TitleName: "贝拉克",
		TitleCode: "b_bellac",
	}
}
