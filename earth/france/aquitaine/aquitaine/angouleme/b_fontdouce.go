package angouleme

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丰杜斯FontdouceBarony struct {
	feud.BaseBarony
}

var BFontdouce丰杜斯 feud.Barony = &丰杜斯FontdouceBarony{}

func init() {
	f := BFontdouce丰杜斯.(*丰杜斯FontdouceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fontdouce",
		TitleName: "丰杜斯",
		TitleCode: "b_fontdouce",
	}
}
