package nedong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雍布拉康YungbulakangBarony struct {
	feud.BaseBarony
}

var BYungbulakang雍布拉康 feud.Barony = &雍布拉康YungbulakangBarony{}

func init() {
	f := BYungbulakang雍布拉康.(*雍布拉康YungbulakangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yungbulakang",
		TitleName: "雍布拉康",
		TitleCode: "b_yungbulakang",
	}
}
