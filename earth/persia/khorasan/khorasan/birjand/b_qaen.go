package birjand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加延QaenBarony struct {
	feud.BaseBarony
}

var BQaen加延 feud.Barony = &加延QaenBarony{}

func init() {
	f := BQaen加延.(*加延QaenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qaen",
		TitleName: "加延",
		TitleCode: "b_qaen",
	}
}
