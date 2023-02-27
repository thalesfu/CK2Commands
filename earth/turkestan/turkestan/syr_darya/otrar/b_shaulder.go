package otrar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绍利杰尔ShaulderBarony struct {
	feud.BaseBarony
}

var BShaulder绍利杰尔 feud.Barony = &绍利杰尔ShaulderBarony{}

func init() {
    f := BShaulder绍利杰尔.(*绍利杰尔ShaulderBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shaulder",
		TitleName: "绍利杰尔",
		TitleCode: "b_shaulder",
	}
}
