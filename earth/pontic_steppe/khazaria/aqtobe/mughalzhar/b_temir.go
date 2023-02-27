package mughalzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 铁米尔TemirBarony struct {
	feud.BaseBarony
}

var BTemir铁米尔 feud.Barony = &铁米尔TemirBarony{}

func init() {
    f := BTemir铁米尔.(*铁米尔TemirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "temir",
		TitleName: "铁米尔",
		TitleCode: "b_temir",
	}
}
