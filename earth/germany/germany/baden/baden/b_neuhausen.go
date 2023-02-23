package baden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺伊豪森NeuhausenBarony struct {
	feud.BaseBarony
}

var BNeuhausen诺伊豪森 feud.Barony = &诺伊豪森NeuhausenBarony{}

func init() {
	f := BNeuhausen诺伊豪森.(*诺伊豪森NeuhausenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neuhausen",
		TitleName: "诺伊豪森",
		TitleCode: "b_neuhausen",
	}
}
