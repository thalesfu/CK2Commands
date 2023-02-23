package byzantion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加拉塔GalataBarony struct {
	feud.BaseBarony
}

var BGalata加拉塔 feud.Barony = &加拉塔GalataBarony{}

func init() {
	f := BGalata加拉塔.(*加拉塔GalataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "galata",
		TitleName: "加拉塔",
		TitleCode: "b_galata",
	}
}
