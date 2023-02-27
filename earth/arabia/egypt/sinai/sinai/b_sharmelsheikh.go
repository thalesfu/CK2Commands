package sinai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙姆沙伊赫SharmelsheikhBarony struct {
	feud.BaseBarony
}

var BSharmelsheikh沙姆沙伊赫 feud.Barony = &沙姆沙伊赫SharmelsheikhBarony{}

func init() {
    f := BSharmelsheikh沙姆沙伊赫.(*沙姆沙伊赫SharmelsheikhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharmelsheikh",
		TitleName: "沙姆沙伊赫",
		TitleCode: "b_sharmelsheikh",
	}
}
