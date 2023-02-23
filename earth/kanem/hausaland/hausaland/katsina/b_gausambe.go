package katsina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈桑布GausambeBarony struct {
	feud.BaseBarony
}

var BGausambe戈桑布 feud.Barony = &戈桑布GausambeBarony{}

func init() {
	f := BGausambe戈桑布.(*戈桑布GausambeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gausambe",
		TitleName: "戈桑布",
		TitleCode: "b_gausambe",
	}
}
