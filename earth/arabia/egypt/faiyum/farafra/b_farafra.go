package farafra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法拉弗拉FarafraBarony struct {
	feud.BaseBarony
}

var BFarafra法拉弗拉 feud.Barony = &法拉弗拉FarafraBarony{}

func init() {
	f := BFarafra法拉弗拉.(*法拉弗拉FarafraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "farafra",
		TitleName: "法拉弗拉",
		TitleCode: "b_farafra",
	}
}
