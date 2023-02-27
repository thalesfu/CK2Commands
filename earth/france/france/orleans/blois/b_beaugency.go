package blois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博让西BeaugencyBarony struct {
	feud.BaseBarony
}

var BBeaugency博让西 feud.Barony = &博让西BeaugencyBarony{}

func init() {
    f := BBeaugency博让西.(*博让西BeaugencyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beaugency",
		TitleName: "博让西",
		TitleCode: "b_beaugency",
	}
}
