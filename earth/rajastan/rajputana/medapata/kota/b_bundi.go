package kota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本迪BundiBarony struct {
	feud.BaseBarony
}

var BBundi本迪 feud.Barony = &本迪BundiBarony{}

func init() {
    f := BBundi本迪.(*本迪BundiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bundi",
		TitleName: "本迪",
		TitleCode: "b_bundi",
	}
}
