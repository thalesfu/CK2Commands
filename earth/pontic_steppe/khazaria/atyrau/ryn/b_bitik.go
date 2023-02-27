package ryn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比蒂克BitikBarony struct {
	feud.BaseBarony
}

var BBitik比蒂克 feud.Barony = &比蒂克BitikBarony{}

func init() {
    f := BBitik比蒂克.(*比蒂克BitikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bitik",
		TitleName: "比蒂克",
		TitleCode: "b_bitik",
	}
}
