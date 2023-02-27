package labourd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克莱朗斯堡LabastideclairenceBarony struct {
	feud.BaseBarony
}

var BLabastideclairence克莱朗斯堡 feud.Barony = &克莱朗斯堡LabastideclairenceBarony{}

func init() {
    f := BLabastideclairence克莱朗斯堡.(*克莱朗斯堡LabastideclairenceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "labastideclairence",
		TitleName: "克莱朗斯堡",
		TitleCode: "b_labastideclairence",
	}
}
