package vashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克巴KebaBarony struct {
	feud.BaseBarony
}

var BKeba克巴 feud.Barony = &克巴KebaBarony{}

func init() {
    f := BKeba克巴.(*克巴KebaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "keba",
		TitleName: "克巴",
		TitleCode: "b_keba",
	}
}
