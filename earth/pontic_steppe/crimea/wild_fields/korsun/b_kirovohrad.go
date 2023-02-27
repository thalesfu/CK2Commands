package korsun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基洛夫格勒KirovohradBarony struct {
	feud.BaseBarony
}

var BKirovohrad基洛夫格勒 feud.Barony = &基洛夫格勒KirovohradBarony{}

func init() {
    f := BKirovohrad基洛夫格勒.(*基洛夫格勒KirovohradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirovohrad",
		TitleName: "基洛夫格勒",
		TitleCode: "b_kirovohrad",
	}
}
