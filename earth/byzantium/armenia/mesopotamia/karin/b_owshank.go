package karin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥乌斯汉克OwshankBarony struct {
	feud.BaseBarony
}

var BOwshank奥乌斯汉克 feud.Barony = &奥乌斯汉克OwshankBarony{}

func init() {
	f := BOwshank奥乌斯汉克.(*奥乌斯汉克OwshankBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "owshank",
		TitleName: "奥乌斯汉克",
		TitleCode: "b_owshank",
	}
}
