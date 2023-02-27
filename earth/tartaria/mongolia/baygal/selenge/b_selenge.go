package selenge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑陵SelengeBarony struct {
	feud.BaseBarony
}

var BSelenge娑陵 feud.Barony = &娑陵SelengeBarony{}

func init() {
    f := BSelenge娑陵.(*娑陵SelengeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "selenge",
		TitleName: "娑陵",
		TitleCode: "b_selenge",
	}
}
