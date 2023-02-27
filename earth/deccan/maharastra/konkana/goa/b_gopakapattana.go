package goa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿波迦钵多那GopakapattanaBarony struct {
	feud.BaseBarony
}

var BGopakapattana瞿波迦钵多那 feud.Barony = &瞿波迦钵多那GopakapattanaBarony{}

func init() {
    f := BGopakapattana瞿波迦钵多那.(*瞿波迦钵多那GopakapattanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gopakapattana",
		TitleName: "瞿波迦钵多那",
		TitleCode: "b_gopakapattana",
	}
}
