package tashkurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 将嘎拉JianggalaBarony struct {
	feud.BaseBarony
}

var BJianggala将嘎拉 feud.Barony = &将嘎拉JianggalaBarony{}

func init() {
    f := BJianggala将嘎拉.(*将嘎拉JianggalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jianggala",
		TitleName: "将嘎拉",
		TitleCode: "b_jianggala",
	}
}
