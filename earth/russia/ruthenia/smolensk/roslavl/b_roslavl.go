package roslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯拉夫尔RoslavlBarony struct {
	feud.BaseBarony
}

var BRoslavl罗斯拉夫尔 feud.Barony = &罗斯拉夫尔RoslavlBarony{}

func init() {
    f := BRoslavl罗斯拉夫尔.(*罗斯拉夫尔RoslavlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roslavl",
		TitleName: "罗斯拉夫尔",
		TitleCode: "b_roslavl",
	}
}
