package tver

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉舍沃BurashevoBarony struct {
	feud.BaseBarony
}

var BBurashevo布拉舍沃 feud.Barony = &布拉舍沃BurashevoBarony{}

func init() {
    f := BBurashevo布拉舍沃.(*布拉舍沃BurashevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burashevo",
		TitleName: "布拉舍沃",
		TitleCode: "b_burashevo",
	}
}
