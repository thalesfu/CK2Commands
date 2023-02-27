package tukums

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎达瓦KandavaBarony struct {
	feud.BaseBarony
}

var BKandava坎达瓦 feud.Barony = &坎达瓦KandavaBarony{}

func init() {
    f := BKandava坎达瓦.(*坎达瓦KandavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kandava",
		TitleName: "坎达瓦",
		TitleCode: "b_kandava",
	}
}
