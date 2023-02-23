package zeila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽拉ZeilaBarony struct {
	feud.BaseBarony
}

var BZeila泽拉 feud.Barony = &泽拉ZeilaBarony{}

func init() {
	f := BZeila泽拉.(*泽拉ZeilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zeila",
		TitleName: "泽拉",
		TitleCode: "b_zeila",
	}
}
