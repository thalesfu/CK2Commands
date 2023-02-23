package leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱昂LeonBarony struct {
	feud.BaseBarony
}

var BLeon莱昂 feud.Barony = &莱昂LeonBarony{}

func init() {
	f := BLeon莱昂.(*莱昂LeonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leon",
		TitleName: "莱昂",
		TitleCode: "b_leon",
	}
}
