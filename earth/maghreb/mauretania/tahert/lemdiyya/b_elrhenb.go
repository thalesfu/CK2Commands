package lemdiyya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾里亨布ElrhenbBarony struct {
	feud.BaseBarony
}

var BElrhenb艾里亨布 feud.Barony = &艾里亨布ElrhenbBarony{}

func init() {
    f := BElrhenb艾里亨布.(*艾里亨布ElrhenbBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elrhenb",
		TitleName: "艾里亨布",
		TitleCode: "b_elrhenb",
	}
}
