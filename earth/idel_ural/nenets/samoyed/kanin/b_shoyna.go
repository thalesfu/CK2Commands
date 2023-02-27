package kanin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绍伊纳ShoynaBarony struct {
	feud.BaseBarony
}

var BShoyna绍伊纳 feud.Barony = &绍伊纳ShoynaBarony{}

func init() {
    f := BShoyna绍伊纳.(*绍伊纳ShoynaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shoyna",
		TitleName: "绍伊纳",
		TitleCode: "b_shoyna",
	}
}
