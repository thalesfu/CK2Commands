package yolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮兹马PizmaBarony struct {
	feud.BaseBarony
}

var BPizma皮兹马 feud.Barony = &皮兹马PizmaBarony{}

func init() {
    f := BPizma皮兹马.(*皮兹马PizmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pizma",
		TitleName: "皮兹马",
		TitleCode: "b_pizma",
	}
}
