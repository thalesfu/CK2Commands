package braganza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗德里戈堡CastelorodrigoBarony struct {
	feud.BaseBarony
}

var BCastelorodrigo罗德里戈堡 feud.Barony = &罗德里戈堡CastelorodrigoBarony{}

func init() {
    f := BCastelorodrigo罗德里戈堡.(*罗德里戈堡CastelorodrigoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castelorodrigo",
		TitleName: "罗德里戈堡",
		TitleCode: "b_castelorodrigo",
	}
}
