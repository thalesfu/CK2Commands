package travunia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博戈罗迪采BogorodiceBarony struct {
	feud.BaseBarony
}

var BBogorodice博戈罗迪采 feud.Barony = &博戈罗迪采BogorodiceBarony{}

func init() {
    f := BBogorodice博戈罗迪采.(*博戈罗迪采BogorodiceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bogorodice",
		TitleName: "博戈罗迪采",
		TitleCode: "b_bogorodice",
	}
}
