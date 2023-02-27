package vijnot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗婆RavelBarony struct {
	feud.BaseBarony
}

var BRavel罗婆 feud.Barony = &罗婆RavelBarony{}

func init() {
    f := BRavel罗婆.(*罗婆RavelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ravel",
		TitleName: "罗婆",
		TitleCode: "b_ravel",
	}
}
