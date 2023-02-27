package dotawo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法拉斯FarasBarony struct {
	feud.BaseBarony
}

var BFaras法拉斯 feud.Barony = &法拉斯FarasBarony{}

func init() {
    f := BFaras法拉斯.(*法拉斯FarasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faras",
		TitleName: "法拉斯",
		TitleCode: "b_faras",
	}
}
