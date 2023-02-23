package vairata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佉住商摩KhatushyamBarony struct {
	feud.BaseBarony
}

var BKhatushyam佉住商摩 feud.Barony = &佉住商摩KhatushyamBarony{}

func init() {
	f := BKhatushyam佉住商摩.(*佉住商摩KhatushyamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khatushyam",
		TitleName: "佉住商摩",
		TitleCode: "b_khatushyam",
	}
}
