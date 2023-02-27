package otuken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 豁里秃马惕Khori_tumedBarony struct {
	feud.BaseBarony
}

var BKhori_tumed豁里秃马惕 feud.Barony = &豁里秃马惕Khori_tumedBarony{}

func init() {
    f := BKhori_tumed豁里秃马惕.(*豁里秃马惕Khori_tumedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khori_tumed",
		TitleName: "豁里秃马惕",
		TitleCode: "b_khori_tumed",
	}
}
