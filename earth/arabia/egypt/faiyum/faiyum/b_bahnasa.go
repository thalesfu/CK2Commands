package faiyum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴赫纳萨BahnasaBarony struct {
	feud.BaseBarony
}

var BBahnasa巴赫纳萨 feud.Barony = &巴赫纳萨BahnasaBarony{}

func init() {
	f := BBahnasa巴赫纳萨.(*巴赫纳萨BahnasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bahnasa",
		TitleName: "巴赫纳萨",
		TitleCode: "b_bahnasa",
	}
}
