package rummah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔雷喀ArraqaiBarony struct {
	feud.BaseBarony
}

var BArraqai阿尔雷喀 feud.Barony = &阿尔雷喀ArraqaiBarony{}

func init() {
    f := BArraqai阿尔雷喀.(*阿尔雷喀ArraqaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arraqai",
		TitleName: "阿尔雷喀",
		TitleCode: "b_arraqai",
	}
}
