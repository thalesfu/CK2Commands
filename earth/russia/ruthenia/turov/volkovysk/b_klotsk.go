package volkovysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克洛茨克KlotskBarony struct {
	feud.BaseBarony
}

var BKlotsk克洛茨克 feud.Barony = &克洛茨克KlotskBarony{}

func init() {
    f := BKlotsk克洛茨克.(*克洛茨克KlotskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klotsk",
		TitleName: "克洛茨克",
		TitleCode: "b_klotsk",
	}
}
