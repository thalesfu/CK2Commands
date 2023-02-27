package ile_de_france

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙福尔拉莫里MontfortlamauryBarony struct {
	feud.BaseBarony
}

var BMontfortlamaury蒙福尔拉莫里 feud.Barony = &蒙福尔拉莫里MontfortlamauryBarony{}

func init() {
    f := BMontfortlamaury蒙福尔拉莫里.(*蒙福尔拉莫里MontfortlamauryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montfortlamaury",
		TitleName: "蒙福尔拉莫里",
		TitleCode: "b_montfortlamaury",
	}
}
