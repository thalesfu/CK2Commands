package koshma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普卢列尔PlurelBarony struct {
	feud.BaseBarony
}

var BPlurel普卢列尔 feud.Barony = &普卢列尔PlurelBarony{}

func init() {
    f := BPlurel普卢列尔.(*普卢列尔PlurelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plurel",
		TitleName: "普卢列尔",
		TitleCode: "b_plurel",
	}
}
