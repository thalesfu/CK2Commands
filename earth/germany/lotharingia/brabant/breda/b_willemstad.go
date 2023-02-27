package breda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 威廉斯塔德WillemstadBarony struct {
	feud.BaseBarony
}

var BWillemstad威廉斯塔德 feud.Barony = &威廉斯塔德WillemstadBarony{}

func init() {
    f := BWillemstad威廉斯塔德.(*威廉斯塔德WillemstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "willemstad",
		TitleName: "威廉斯塔德",
		TitleCode: "b_willemstad",
	}
}
