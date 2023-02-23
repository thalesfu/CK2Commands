package lolland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班霍尔姆BandholmBarony struct {
	feud.BaseBarony
}

var BBandholm班霍尔姆 feud.Barony = &班霍尔姆BandholmBarony{}

func init() {
	f := BBandholm班霍尔姆.(*班霍尔姆BandholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bandholm",
		TitleName: "班霍尔姆",
		TitleCode: "b_bandholm",
	}
}
