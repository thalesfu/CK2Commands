package alania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃夫努什基VovnushkiBarony struct {
	feud.BaseBarony
}

var BVovnushki沃夫努什基 feud.Barony = &沃夫努什基VovnushkiBarony{}

func init() {
    f := BVovnushki沃夫努什基.(*沃夫努什基VovnushkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vovnushki",
		TitleName: "沃夫努什基",
		TitleCode: "b_vovnushki",
	}
}
