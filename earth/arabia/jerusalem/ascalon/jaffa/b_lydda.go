package jaffa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕大LyddaBarony struct {
	feud.BaseBarony
}

var BLydda吕大 feud.Barony = &吕大LyddaBarony{}

func init() {
    f := BLydda吕大.(*吕大LyddaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lydda",
		TitleName: "吕大",
		TitleCode: "b_lydda",
	}
}
