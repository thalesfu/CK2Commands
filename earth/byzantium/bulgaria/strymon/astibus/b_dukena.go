package astibus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜凯纳DukenaBarony struct {
	feud.BaseBarony
}

var BDukena杜凯纳 feud.Barony = &杜凯纳DukenaBarony{}

func init() {
    f := BDukena杜凯纳.(*杜凯纳DukenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dukena",
		TitleName: "杜凯纳",
		TitleCode: "b_dukena",
	}
}
