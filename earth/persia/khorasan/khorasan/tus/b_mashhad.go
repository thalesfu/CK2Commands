package tus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马什哈德MashhadBarony struct {
	feud.BaseBarony
}

var BMashhad马什哈德 feud.Barony = &马什哈德MashhadBarony{}

func init() {
    f := BMashhad马什哈德.(*马什哈德MashhadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mashhad",
		TitleName: "马什哈德",
		TitleCode: "b_mashhad",
	}
}
