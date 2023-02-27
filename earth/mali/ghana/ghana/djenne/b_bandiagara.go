package djenne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邦贾加拉BandiagaraBarony struct {
	feud.BaseBarony
}

var BBandiagara邦贾加拉 feud.Barony = &邦贾加拉BandiagaraBarony{}

func init() {
    f := BBandiagara邦贾加拉.(*邦贾加拉BandiagaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bandiagara",
		TitleName: "邦贾加拉",
		TitleCode: "b_bandiagara",
	}
}
