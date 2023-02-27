package al_bichri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈亚丁MayadinBarony struct {
	feud.BaseBarony
}

var BMayadin迈亚丁 feud.Barony = &迈亚丁MayadinBarony{}

func init() {
    f := BMayadin迈亚丁.(*迈亚丁MayadinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mayadin",
		TitleName: "迈亚丁",
		TitleCode: "b_mayadin",
	}
}
