package gower

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉赫尔LoughorBarony struct {
	feud.BaseBarony
}

var BLoughor拉赫尔 feud.Barony = &拉赫尔LoughorBarony{}

func init() {
    f := BLoughor拉赫尔.(*拉赫尔LoughorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loughor",
		TitleName: "拉赫尔",
		TitleCode: "b_loughor",
	}
}
