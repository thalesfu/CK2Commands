package dakhina_desa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 屈胝KotteBarony struct {
	feud.BaseBarony
}

var BKotte屈胝 feud.Barony = &屈胝KotteBarony{}

func init() {
    f := BKotte屈胝.(*屈胝KotteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotte",
		TitleName: "屈胝",
		TitleCode: "b_kotte",
	}
}
