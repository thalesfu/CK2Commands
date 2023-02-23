package chikoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘多特沃尔察ChudotvortsaBarony struct {
	feud.BaseBarony
}

var BChudotvortsa丘多特沃尔察 feud.Barony = &丘多特沃尔察ChudotvortsaBarony{}

func init() {
	f := BChudotvortsa丘多特沃尔察.(*丘多特沃尔察ChudotvortsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chudotvortsa",
		TitleName: "丘多特沃尔察",
		TitleCode: "b_chudotvortsa",
	}
}
