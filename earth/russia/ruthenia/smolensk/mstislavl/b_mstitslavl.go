package mstislavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆斯季斯拉夫尔MstitslavlBarony struct {
	feud.BaseBarony
}

var BMstitslavl姆斯季斯拉夫尔 feud.Barony = &姆斯季斯拉夫尔MstitslavlBarony{}

func init() {
    f := BMstitslavl姆斯季斯拉夫尔.(*姆斯季斯拉夫尔MstitslavlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mstitslavl",
		TitleName: "姆斯季斯拉夫尔",
		TitleCode: "b_mstitslavl",
	}
}
