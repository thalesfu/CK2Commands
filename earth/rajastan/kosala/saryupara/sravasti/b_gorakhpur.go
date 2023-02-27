package sravasti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿罗叉城GorakhpurBarony struct {
	feud.BaseBarony
}

var BGorakhpur瞿罗叉城 feud.Barony = &瞿罗叉城GorakhpurBarony{}

func init() {
    f := BGorakhpur瞿罗叉城.(*瞿罗叉城GorakhpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorakhpur",
		TitleName: "瞿罗叉城",
		TitleCode: "b_gorakhpur",
	}
}
