package yaroslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅罗斯拉夫尔YaroslavlBarony struct {
	feud.BaseBarony
}

var BYaroslavl雅罗斯拉夫尔 feud.Barony = &雅罗斯拉夫尔YaroslavlBarony{}

func init() {
	f := BYaroslavl雅罗斯拉夫尔.(*雅罗斯拉夫尔YaroslavlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yaroslavl",
		TitleName: "雅罗斯拉夫尔",
		TitleCode: "b_yaroslavl",
	}
}
