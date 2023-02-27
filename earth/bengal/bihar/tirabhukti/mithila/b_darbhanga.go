package mithila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达罗傍伽DarbhangaBarony struct {
	feud.BaseBarony
}

var BDarbhanga达罗傍伽 feud.Barony = &达罗傍伽DarbhangaBarony{}

func init() {
    f := BDarbhanga达罗傍伽.(*达罗傍伽DarbhangaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darbhanga",
		TitleName: "达罗傍伽",
		TitleCode: "b_darbhanga",
	}
}
