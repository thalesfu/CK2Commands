package kota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘吒KotaBarony struct {
	feud.BaseBarony
}

var BKota拘吒 feud.Barony = &拘吒KotaBarony{}

func init() {
    f := BKota拘吒.(*拘吒KotaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kota",
		TitleName: "拘吒",
		TitleCode: "b_kota",
	}
}
