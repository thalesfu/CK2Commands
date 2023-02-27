package saintois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑图瓦SaintoisBarony struct {
	feud.BaseBarony
}

var BSaintois桑图瓦 feud.Barony = &桑图瓦SaintoisBarony{}

func init() {
    f := BSaintois桑图瓦.(*桑图瓦SaintoisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintois",
		TitleName: "桑图瓦",
		TitleCode: "b_saintois",
	}
}
