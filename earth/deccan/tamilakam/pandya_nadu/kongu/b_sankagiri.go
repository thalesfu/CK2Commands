package kongu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 商迦耆厘SankagiriBarony struct {
	feud.BaseBarony
}

var BSankagiri商迦耆厘 feud.Barony = &商迦耆厘SankagiriBarony{}

func init() {
    f := BSankagiri商迦耆厘.(*商迦耆厘SankagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sankagiri",
		TitleName: "商迦耆厘",
		TitleCode: "b_sankagiri",
	}
}
