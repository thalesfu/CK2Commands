package kolomna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯季斯拉夫尔RostislavlBarony struct {
	feud.BaseBarony
}

var BRostislavl罗斯季斯拉夫尔 feud.Barony = &罗斯季斯拉夫尔RostislavlBarony{}

func init() {
    f := BRostislavl罗斯季斯拉夫尔.(*罗斯季斯拉夫尔RostislavlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rostislavl",
		TitleName: "罗斯季斯拉夫尔",
		TitleCode: "b_rostislavl",
	}
}
