package gezira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科伊贾姆KoijamBarony struct {
	feud.BaseBarony
}

var BKoijam科伊贾姆 feud.Barony = &科伊贾姆KoijamBarony{}

func init() {
    f := BKoijam科伊贾姆.(*科伊贾姆KoijamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koijam",
		TitleName: "科伊贾姆",
		TitleCode: "b_koijam",
	}
}
