package faereyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卡尼锡SkansinBarony struct {
	feud.BaseBarony
}

var BSkansin斯卡尼锡 feud.Barony = &斯卡尼锡SkansinBarony{}

func init() {
    f := BSkansin斯卡尼锡.(*斯卡尼锡SkansinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skansin",
		TitleName: "斯卡尼锡",
		TitleCode: "b_skansin",
	}
}
