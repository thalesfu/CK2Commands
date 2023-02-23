package orania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马格尼耶MaghniaBarony struct {
	feud.BaseBarony
}

var BMaghnia马格尼耶 feud.Barony = &马格尼耶MaghniaBarony{}

func init() {
	f := BMaghnia马格尼耶.(*马格尼耶MaghniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maghnia",
		TitleName: "马格尼耶",
		TitleCode: "b_maghnia",
	}
}
