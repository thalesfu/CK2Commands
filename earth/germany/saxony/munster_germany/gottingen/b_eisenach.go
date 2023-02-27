package gottingen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾森纳赫EisenachBarony struct {
	feud.BaseBarony
}

var BEisenach艾森纳赫 feud.Barony = &艾森纳赫EisenachBarony{}

func init() {
    f := BEisenach艾森纳赫.(*艾森纳赫EisenachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eisenach",
		TitleName: "艾森纳赫",
		TitleCode: "b_eisenach",
	}
}
