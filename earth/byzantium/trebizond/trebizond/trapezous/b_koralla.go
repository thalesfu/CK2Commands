package trapezous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍雷拉KorallaBarony struct {
	feud.BaseBarony
}

var BKoralla霍雷拉 feud.Barony = &霍雷拉KorallaBarony{}

func init() {
	f := BKoralla霍雷拉.(*霍雷拉KorallaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koralla",
		TitleName: "霍雷拉",
		TitleCode: "b_koralla",
	}
}
