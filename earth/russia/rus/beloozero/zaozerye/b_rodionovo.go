package zaozerye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗季奥诺沃RodionovoBarony struct {
	feud.BaseBarony
}

var BRodionovo罗季奥诺沃 feud.Barony = &罗季奥诺沃RodionovoBarony{}

func init() {
    f := BRodionovo罗季奥诺沃.(*罗季奥诺沃RodionovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rodionovo",
		TitleName: "罗季奥诺沃",
		TitleCode: "b_rodionovo",
	}
}
