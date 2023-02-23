package khantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊格里姆IgrimBarony struct {
	feud.BaseBarony
}

var BIgrim伊格里姆 feud.Barony = &伊格里姆IgrimBarony{}

func init() {
	f := BIgrim伊格里姆.(*伊格里姆IgrimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "igrim",
		TitleName: "伊格里姆",
		TitleCode: "b_igrim",
	}
}
