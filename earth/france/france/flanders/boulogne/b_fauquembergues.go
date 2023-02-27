package boulogne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福康贝格FauquemberguesBarony struct {
	feud.BaseBarony
}

var BFauquembergues福康贝格 feud.Barony = &福康贝格FauquemberguesBarony{}

func init() {
    f := BFauquembergues福康贝格.(*福康贝格FauquemberguesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fauquembergues",
		TitleName: "福康贝格",
		TitleCode: "b_fauquembergues",
	}
}
