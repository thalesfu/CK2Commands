package urbino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法诺FanoBarony struct {
	feud.BaseBarony
}

var BFano法诺 feud.Barony = &法诺FanoBarony{}

func init() {
	f := BFano法诺.(*法诺FanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fano",
		TitleName: "法诺",
		TitleCode: "b_fano",
	}
}
