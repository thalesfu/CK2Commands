package amol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尔扎别克MirzabekBarony struct {
	feud.BaseBarony
}

var BMirzabek米尔扎别克 feud.Barony = &米尔扎别克MirzabekBarony{}

func init() {
    f := BMirzabek米尔扎别克.(*米尔扎别克MirzabekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirzabek",
		TitleName: "米尔扎别克",
		TitleCode: "b_mirzabek",
	}
}
