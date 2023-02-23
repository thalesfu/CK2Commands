package asayita

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿凡博AfamboBarony struct {
	feud.BaseBarony
}

var BAfambo阿凡博 feud.Barony = &阿凡博AfamboBarony{}

func init() {
	f := BAfambo阿凡博.(*阿凡博AfamboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "afambo",
		TitleName: "阿凡博",
		TitleCode: "b_afambo",
	}
}
