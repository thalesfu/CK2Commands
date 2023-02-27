package mantua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼图阿MantuaBarony struct {
	feud.BaseBarony
}

var BMantua曼图阿 feud.Barony = &曼图阿MantuaBarony{}

func init() {
    f := BMantua曼图阿.(*曼图阿MantuaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mantua",
		TitleName: "曼图阿",
		TitleCode: "b_mantua",
	}
}
