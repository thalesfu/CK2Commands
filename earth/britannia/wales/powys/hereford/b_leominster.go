package hereford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱姆斯特LeominsterBarony struct {
	feud.BaseBarony
}

var BLeominster莱姆斯特 feud.Barony = &莱姆斯特LeominsterBarony{}

func init() {
	f := BLeominster莱姆斯特.(*莱姆斯特LeominsterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leominster",
		TitleName: "莱姆斯特",
		TitleCode: "b_leominster",
	}
}
