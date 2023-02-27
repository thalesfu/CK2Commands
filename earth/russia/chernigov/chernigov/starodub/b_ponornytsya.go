package starodub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波诺尔尼察PonornytsyaBarony struct {
	feud.BaseBarony
}

var BPonornytsya波诺尔尼察 feud.Barony = &波诺尔尼察PonornytsyaBarony{}

func init() {
    f := BPonornytsya波诺尔尼察.(*波诺尔尼察PonornytsyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ponornytsya",
		TitleName: "波诺尔尼察",
		TitleCode: "b_ponornytsya",
	}
}
