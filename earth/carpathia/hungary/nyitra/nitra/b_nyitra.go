package nitra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼特劳NyitraBarony struct {
	feud.BaseBarony
}

var BNyitra尼特劳 feud.Barony = &尼特劳NyitraBarony{}

func init() {
	f := BNyitra尼特劳.(*尼特劳NyitraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyitra",
		TitleName: "尼特劳",
		TitleCode: "b_nyitra",
	}
}
