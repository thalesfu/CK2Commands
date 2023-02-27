package armail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴斯蒂巴茜拉Basti_basiraBarony struct {
	feud.BaseBarony
}

var BBasti_basira巴斯蒂巴茜拉 feud.Barony = &巴斯蒂巴茜拉Basti_basiraBarony{}

func init() {
    f := BBasti_basira巴斯蒂巴茜拉.(*巴斯蒂巴茜拉Basti_basiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "basti_basira",
		TitleName: "巴斯蒂巴茜拉",
		TitleCode: "b_basti_basira",
	}
}
