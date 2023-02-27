package gower

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯旺西SwanseaBarony struct {
	feud.BaseBarony
}

var BSwansea斯旺西 feud.Barony = &斯旺西SwanseaBarony{}

func init() {
    f := BSwansea斯旺西.(*斯旺西SwanseaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "swansea",
		TitleName: "斯旺西",
		TitleCode: "b_swansea",
	}
}
