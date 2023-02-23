package infa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡里卜盖KhouribgaBarony struct {
	feud.BaseBarony
}

var BKhouribga胡里卜盖 feud.Barony = &胡里卜盖KhouribgaBarony{}

func init() {
	f := BKhouribga胡里卜盖.(*胡里卜盖KhouribgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khouribga",
		TitleName: "胡里卜盖",
		TitleCode: "b_khouribga",
	}
}
