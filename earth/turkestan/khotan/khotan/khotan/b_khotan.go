package khotan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于阗KhotanBarony struct {
	feud.BaseBarony
}

var BKhotan于阗 feud.Barony = &于阗KhotanBarony{}

func init() {
    f := BKhotan于阗.(*于阗KhotanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khotan",
		TitleName: "于阗",
		TitleCode: "b_khotan",
	}
}
