package kuma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基兹利亚尔KizlyarBarony struct {
	feud.BaseBarony
}

var BKizlyar基兹利亚尔 feud.Barony = &基兹利亚尔KizlyarBarony{}

func init() {
    f := BKizlyar基兹利亚尔.(*基兹利亚尔KizlyarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kizlyar",
		TitleName: "基兹利亚尔",
		TitleCode: "b_kizlyar",
	}
}
