package mertola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫朗MouraoBarony struct {
	feud.BaseBarony
}

var BMourao莫朗 feud.Barony = &莫朗MouraoBarony{}

func init() {
    f := BMourao莫朗.(*莫朗MouraoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mourao",
		TitleName: "莫朗",
		TitleCode: "b_mourao",
	}
}
