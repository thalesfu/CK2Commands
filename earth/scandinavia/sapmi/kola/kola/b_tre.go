package kola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷TreBarony struct {
	feud.BaseBarony
}

var BTre特雷 feud.Barony = &特雷TreBarony{}

func init() {
    f := BTre特雷.(*特雷TreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tre",
		TitleName: "特雷",
		TitleCode: "b_tre",
	}
}
