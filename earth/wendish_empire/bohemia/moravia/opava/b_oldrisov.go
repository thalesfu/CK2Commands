package opava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔德日绍夫OldrisovBarony struct {
	feud.BaseBarony
}

var BOldrisov奥尔德日绍夫 feud.Barony = &奥尔德日绍夫OldrisovBarony{}

func init() {
    f := BOldrisov奥尔德日绍夫.(*奥尔德日绍夫OldrisovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oldrisov",
		TitleName: "奥尔德日绍夫",
		TitleCode: "b_oldrisov",
	}
}
