package uch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邬脂UchBarony struct {
	feud.BaseBarony
}

var BUch邬脂 feud.Barony = &邬脂UchBarony{}

func init() {
    f := BUch邬脂.(*邬脂UchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uch",
		TitleName: "邬脂",
		TitleCode: "b_uch",
	}
}
