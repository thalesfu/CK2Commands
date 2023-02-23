package constantine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提吉斯TijisBarony struct {
	feud.BaseBarony
}

var BTijis提吉斯 feud.Barony = &提吉斯TijisBarony{}

func init() {
	f := BTijis提吉斯.(*提吉斯TijisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tijis",
		TitleName: "提吉斯",
		TitleCode: "b_tijis",
	}
}
