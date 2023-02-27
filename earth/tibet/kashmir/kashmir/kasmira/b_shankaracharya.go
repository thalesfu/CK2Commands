package kasmira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 商羯罗阿遮利耶ShankaracharyaBarony struct {
	feud.BaseBarony
}

var BShankaracharya商羯罗阿遮利耶 feud.Barony = &商羯罗阿遮利耶ShankaracharyaBarony{}

func init() {
    f := BShankaracharya商羯罗阿遮利耶.(*商羯罗阿遮利耶ShankaracharyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shankaracharya",
		TitleName: "商羯罗阿遮利耶",
		TitleCode: "b_shankaracharya",
	}
}
