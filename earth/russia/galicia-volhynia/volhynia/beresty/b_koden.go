package beresty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科登KodenBarony struct {
	feud.BaseBarony
}

var BKoden科登 feud.Barony = &科登KodenBarony{}

func init() {
    f := BKoden科登.(*科登KodenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koden",
		TitleName: "科登",
		TitleCode: "b_koden",
	}
}
