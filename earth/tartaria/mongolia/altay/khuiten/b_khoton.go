package khuiten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 和屯KhotonBarony struct {
	feud.BaseBarony
}

var BKhoton和屯 feud.Barony = &和屯KhotonBarony{}

func init() {
    f := BKhoton和屯.(*和屯KhotonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khoton",
		TitleName: "和屯",
		TitleCode: "b_khoton",
	}
}
