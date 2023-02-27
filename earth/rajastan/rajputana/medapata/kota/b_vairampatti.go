package kota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗蓝钵胝VairampattiBarony struct {
	feud.BaseBarony
}

var BVairampatti毗蓝钵胝 feud.Barony = &毗蓝钵胝VairampattiBarony{}

func init() {
    f := BVairampatti毗蓝钵胝.(*毗蓝钵胝VairampattiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vairampatti",
		TitleName: "毗蓝钵胝",
		TitleCode: "b_vairampatti",
	}
}
