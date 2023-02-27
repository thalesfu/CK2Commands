package ascalon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫比贾HarbijahBarony struct {
	feud.BaseBarony
}

var BHarbijah赫比贾 feud.Barony = &赫比贾HarbijahBarony{}

func init() {
    f := BHarbijah赫比贾.(*赫比贾HarbijahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harbijah",
		TitleName: "赫比贾",
		TitleCode: "b_harbijah",
	}
}
