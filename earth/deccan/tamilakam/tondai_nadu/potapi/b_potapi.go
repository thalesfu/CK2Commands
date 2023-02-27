package potapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩多毗PotapiBarony struct {
	feud.BaseBarony
}

var BPotapi菩多毗 feud.Barony = &菩多毗PotapiBarony{}

func init() {
    f := BPotapi菩多毗.(*菩多毗PotapiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "potapi",
		TitleName: "菩多毗",
		TitleCode: "b_potapi",
	}
}
