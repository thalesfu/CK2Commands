package asturias_de_oviedo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎加斯德奥尼斯CangasdeonisBarony struct {
	feud.BaseBarony
}

var BCangasdeonis坎加斯德奥尼斯 feud.Barony = &坎加斯德奥尼斯CangasdeonisBarony{}

func init() {
    f := BCangasdeonis坎加斯德奥尼斯.(*坎加斯德奥尼斯CangasdeonisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cangasdeonis",
		TitleName: "坎加斯德奥尼斯",
		TitleCode: "b_cangasdeonis",
	}
}
