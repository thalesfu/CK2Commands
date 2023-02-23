package rovaniemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧纳斯OunasBarony struct {
	feud.BaseBarony
}

var BOunas欧纳斯 feud.Barony = &欧纳斯OunasBarony{}

func init() {
	f := BOunas欧纳斯.(*欧纳斯OunasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ounas",
		TitleName: "欧纳斯",
		TitleCode: "b_ounas",
	}
}
