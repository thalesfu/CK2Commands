package leiningen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴滕贝格BattenbergBarony struct {
	feud.BaseBarony
}

var BBattenberg巴滕贝格 feud.Barony = &巴滕贝格BattenbergBarony{}

func init() {
	f := BBattenberg巴滕贝格.(*巴滕贝格BattenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "battenberg",
		TitleName: "巴滕贝格",
		TitleCode: "b_battenberg",
	}
}
