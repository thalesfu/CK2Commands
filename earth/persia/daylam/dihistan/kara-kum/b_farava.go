package kara-kum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法拉瓦FaravaBarony struct {
	feud.BaseBarony
}

var BFarava法拉瓦 feud.Barony = &法拉瓦FaravaBarony{}

func init() {
    f := BFarava法拉瓦.(*法拉瓦FaravaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "farava",
		TitleName: "法拉瓦",
		TitleCode: "b_farava",
	}
}
