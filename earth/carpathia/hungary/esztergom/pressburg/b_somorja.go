package pressburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绍莫尔尧SomorjaBarony struct {
	feud.BaseBarony
}

var BSomorja绍莫尔尧 feud.Barony = &绍莫尔尧SomorjaBarony{}

func init() {
    f := BSomorja绍莫尔尧.(*绍莫尔尧SomorjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "somorja",
		TitleName: "绍莫尔尧",
		TitleCode: "b_somorja",
	}
}
