package dublin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 芬格拉斯FinglasBarony struct {
	feud.BaseBarony
}

var BFinglas芬格拉斯 feud.Barony = &芬格拉斯FinglasBarony{}

func init() {
	f := BFinglas芬格拉斯.(*芬格拉斯FinglasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "finglas",
		TitleName: "芬格拉斯",
		TitleCode: "b_finglas",
	}
}
