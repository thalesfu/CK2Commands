package gobir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比尔尼恩孔尼Birnin_konniBarony struct {
	feud.BaseBarony
}

var BBirnin_konni比尔尼恩孔尼 feud.Barony = &比尔尼恩孔尼Birnin_konniBarony{}

func init() {
    f := BBirnin_konni比尔尼恩孔尼.(*比尔尼恩孔尼Birnin_konniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birnin_konni",
		TitleName: "比尔尼恩孔尼",
		TitleCode: "b_birnin_konni",
	}
}
