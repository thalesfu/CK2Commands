package pratishthana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾哈迈德纳格尔AhmednagarBarony struct {
	feud.BaseBarony
}

var BAhmednagar艾哈迈德纳格尔 feud.Barony = &艾哈迈德纳格尔AhmednagarBarony{}

func init() {
	f := BAhmednagar艾哈迈德纳格尔.(*艾哈迈德纳格尔AhmednagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ahmednagar",
		TitleName: "艾哈迈德纳格尔",
		TitleCode: "b_ahmednagar",
	}
}
