package yuni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 石坑ShikengBarony struct {
	feud.BaseBarony
}

var BShikeng石坑 feud.Barony = &石坑ShikengBarony{}

func init() {
	f := BShikeng石坑.(*石坑ShikengBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shikeng",
		TitleName: "石坑",
		TitleCode: "b_shikeng",
	}
}
