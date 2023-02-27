package limisol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科洛西KolossiBarony struct {
	feud.BaseBarony
}

var BKolossi科洛西 feud.Barony = &科洛西KolossiBarony{}

func init() {
    f := BKolossi科洛西.(*科洛西KolossiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolossi",
		TitleName: "科洛西",
		TitleCode: "b_kolossi",
	}
}
