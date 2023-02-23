package imotski

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利夫诺LivnoBarony struct {
	feud.BaseBarony
}

var BLivno利夫诺 feud.Barony = &利夫诺LivnoBarony{}

func init() {
	f := BLivno利夫诺.(*利夫诺LivnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "livno",
		TitleName: "利夫诺",
		TitleCode: "b_livno",
	}
}
