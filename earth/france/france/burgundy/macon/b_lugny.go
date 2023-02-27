package macon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕尼LugnyBarony struct {
	feud.BaseBarony
}

var BLugny吕尼 feud.Barony = &吕尼LugnyBarony{}

func init() {
    f := BLugny吕尼.(*吕尼LugnyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lugny",
		TitleName: "吕尼",
		TitleCode: "b_lugny",
	}
}
