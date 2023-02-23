package mangyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宗嘎DzongkaBarony struct {
	feud.BaseBarony
}

var BDzongka宗嘎 feud.Barony = &宗嘎DzongkaBarony{}

func init() {
	f := BDzongka宗嘎.(*宗嘎DzongkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dzongka",
		TitleName: "宗嘎",
		TitleCode: "b_dzongka",
	}
}
