package hebron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希伯伦HebronBarony struct {
	feud.BaseBarony
}

var BHebron希伯伦 feud.Barony = &希伯伦HebronBarony{}

func init() {
    f := BHebron希伯伦.(*希伯伦HebronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hebron",
		TitleName: "希伯伦",
		TitleCode: "b_hebron",
	}
}
