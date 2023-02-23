package napata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科赫伊利KeheilliBarony struct {
	feud.BaseBarony
}

var BKeheilli科赫伊利 feud.Barony = &科赫伊利KeheilliBarony{}

func init() {
	f := BKeheilli科赫伊利.(*科赫伊利KeheilliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "keheilli",
		TitleName: "科赫伊利",
		TitleCode: "b_keheilli",
	}
}
