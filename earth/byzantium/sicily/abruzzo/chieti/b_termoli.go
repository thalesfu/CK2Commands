package chieti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰尔莫利TermoliBarony struct {
	feud.BaseBarony
}

var BTermoli泰尔莫利 feud.Barony = &泰尔莫利TermoliBarony{}

func init() {
    f := BTermoli泰尔莫利.(*泰尔莫利TermoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "termoli",
		TitleName: "泰尔莫利",
		TitleCode: "b_termoli",
	}
}
