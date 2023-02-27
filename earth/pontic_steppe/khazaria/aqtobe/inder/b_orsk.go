package inder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因德博尔OrskBarony struct {
	feud.BaseBarony
}

var BOrsk因德博尔 feud.Barony = &因德博尔OrskBarony{}

func init() {
    f := BOrsk因德博尔.(*因德博尔OrskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orsk",
		TitleName: "因德博尔",
		TitleCode: "b_orsk",
	}
}
