package sundgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔马KolmarBarony struct {
	feud.BaseBarony
}

var BKolmar科尔马 feud.Barony = &科尔马KolmarBarony{}

func init() {
	f := BKolmar科尔马.(*科尔马KolmarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolmar",
		TitleName: "科尔马",
		TitleCode: "b_kolmar",
	}
}
