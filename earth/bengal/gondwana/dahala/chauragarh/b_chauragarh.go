package chauragarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 招罗迦罗ChauragarhBarony struct {
	feud.BaseBarony
}

var BChauragarh招罗迦罗 feud.Barony = &招罗迦罗ChauragarhBarony{}

func init() {
    f := BChauragarh招罗迦罗.(*招罗迦罗ChauragarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chauragarh",
		TitleName: "招罗迦罗",
		TitleCode: "b_chauragarh",
	}
}
