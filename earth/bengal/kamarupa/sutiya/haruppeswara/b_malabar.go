package haruppeswara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麻啰拔MalabarBarony struct {
	feud.BaseBarony
}

var BMalabar麻啰拔 feud.Barony = &麻啰拔MalabarBarony{}

func init() {
    f := BMalabar麻啰拔.(*麻啰拔MalabarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malabar",
		TitleName: "麻啰拔",
		TitleCode: "b_malabar",
	}
}
