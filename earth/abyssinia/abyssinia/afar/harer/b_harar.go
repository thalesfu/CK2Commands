package harer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈勒尔HararBarony struct {
	feud.BaseBarony
}

var BHarar哈勒尔 feud.Barony = &哈勒尔HararBarony{}

func init() {
    f := BHarar哈勒尔.(*哈勒尔HararBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harar",
		TitleName: "哈勒尔",
		TitleCode: "b_harar",
	}
}
