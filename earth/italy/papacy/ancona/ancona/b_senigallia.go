package ancona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尼加利亚SenigalliaBarony struct {
	feud.BaseBarony
}

var BSenigallia塞尼加利亚 feud.Barony = &塞尼加利亚SenigalliaBarony{}

func init() {
	f := BSenigallia塞尼加利亚.(*塞尼加利亚SenigalliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "senigallia",
		TitleName: "塞尼加利亚",
		TitleCode: "b_senigallia",
	}
}
