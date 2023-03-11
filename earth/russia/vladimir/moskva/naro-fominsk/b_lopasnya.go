package naro-fominsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛帕斯尼亚LopasnyaBarony struct {
	feud.BaseBarony
}

var BLopasnya洛帕斯尼亚 feud.Barony = &洛帕斯尼亚LopasnyaBarony{}

func init() {
    f := BLopasnya洛帕斯尼亚.(*洛帕斯尼亚LopasnyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lopasnya",
		TitleName: "洛帕斯尼亚",
		TitleCode: "b_lopasnya",
	}
}
