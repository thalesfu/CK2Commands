package vozviahel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科列茨KoretsBarony struct {
	feud.BaseBarony
}

var BKorets科列茨 feud.Barony = &科列茨KoretsBarony{}

func init() {
	f := BKorets科列茨.(*科列茨KoretsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korets",
		TitleName: "科列茨",
		TitleCode: "b_korets",
	}
}
