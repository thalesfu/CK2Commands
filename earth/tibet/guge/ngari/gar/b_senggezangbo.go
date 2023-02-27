package gar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 森格藏布SenggezangboBarony struct {
	feud.BaseBarony
}

var BSenggezangbo森格藏布 feud.Barony = &森格藏布SenggezangboBarony{}

func init() {
    f := BSenggezangbo森格藏布.(*森格藏布SenggezangboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "senggezangbo",
		TitleName: "森格藏布",
		TitleCode: "b_senggezangbo",
	}
}
