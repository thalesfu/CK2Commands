package khovsgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔尚ArshanBarony struct {
	feud.BaseBarony
}

var BArshan阿尔尚 feud.Barony = &阿尔尚ArshanBarony{}

func init() {
	f := BArshan阿尔尚.(*阿尔尚ArshanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arshan",
		TitleName: "阿尔尚",
		TitleCode: "b_arshan",
	}
}
