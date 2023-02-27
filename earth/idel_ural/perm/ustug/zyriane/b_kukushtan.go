package zyriane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库库什坦KukushtanBarony struct {
	feud.BaseBarony
}

var BKukushtan库库什坦 feud.Barony = &库库什坦KukushtanBarony{}

func init() {
    f := BKukushtan库库什坦.(*库库什坦KukushtanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kukushtan",
		TitleName: "库库什坦",
		TitleCode: "b_kukushtan",
	}
}
