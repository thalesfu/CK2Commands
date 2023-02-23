package bari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比通托BitontoBarony struct {
	feud.BaseBarony
}

var BBitonto比通托 feud.Barony = &比通托BitontoBarony{}

func init() {
	f := BBitonto比通托.(*比通托BitontoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bitonto",
		TitleName: "比通托",
		TitleCode: "b_bitonto",
	}
}
