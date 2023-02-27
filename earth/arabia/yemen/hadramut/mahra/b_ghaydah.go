package mahra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖达GhaydahBarony struct {
	feud.BaseBarony
}

var BGhaydah盖达 feud.Barony = &盖达GhaydahBarony{}

func init() {
    f := BGhaydah盖达.(*盖达GhaydahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghaydah",
		TitleName: "盖达",
		TitleCode: "b_ghaydah",
	}
}
