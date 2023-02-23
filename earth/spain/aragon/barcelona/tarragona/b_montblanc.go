package tarragona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙布朗MontblancBarony struct {
	feud.BaseBarony
}

var BMontblanc蒙布朗 feud.Barony = &蒙布朗MontblancBarony{}

func init() {
	f := BMontblanc蒙布朗.(*蒙布朗MontblancBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montblanc",
		TitleName: "蒙布朗",
		TitleCode: "b_montblanc",
	}
}
