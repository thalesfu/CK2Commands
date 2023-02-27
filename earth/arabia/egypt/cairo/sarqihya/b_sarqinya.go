package sarqihya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比勒拜斯SarqinyaBarony struct {
	feud.BaseBarony
}

var BSarqinya比勒拜斯 feud.Barony = &比勒拜斯SarqinyaBarony{}

func init() {
    f := BSarqinya比勒拜斯.(*比勒拜斯SarqinyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarqinya",
		TitleName: "比勒拜斯",
		TitleCode: "b_sarqinya",
	}
}
