package amastris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提乌姆TiumBarony struct {
	feud.BaseBarony
}

var BTium提乌姆 feud.Barony = &提乌姆TiumBarony{}

func init() {
    f := BTium提乌姆.(*提乌姆TiumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tium",
		TitleName: "提乌姆",
		TitleCode: "b_tium",
	}
}
