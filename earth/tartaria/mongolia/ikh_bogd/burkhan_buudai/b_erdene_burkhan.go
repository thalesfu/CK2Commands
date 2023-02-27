package burkhan_buudai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 额尔德尼Erdene_burkhanBarony struct {
	feud.BaseBarony
}

var BErdene_burkhan额尔德尼 feud.Barony = &额尔德尼Erdene_burkhanBarony{}

func init() {
    f := BErdene_burkhan额尔德尼.(*额尔德尼Erdene_burkhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erdene_burkhan",
		TitleName: "额尔德尼",
		TitleCode: "b_erdene_burkhan",
	}
}
