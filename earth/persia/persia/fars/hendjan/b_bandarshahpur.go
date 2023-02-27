package hendjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班达尔沙普尔BandarshahpurBarony struct {
	feud.BaseBarony
}

var BBandarshahpur班达尔沙普尔 feud.Barony = &班达尔沙普尔BandarshahpurBarony{}

func init() {
    f := BBandarshahpur班达尔沙普尔.(*班达尔沙普尔BandarshahpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bandarshahpur",
		TitleName: "班达尔沙普尔",
		TitleCode: "b_bandarshahpur",
	}
}
