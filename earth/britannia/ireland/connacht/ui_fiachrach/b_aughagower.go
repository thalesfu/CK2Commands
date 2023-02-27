package ui_fiachrach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿哈乔韦尔AughagowerBarony struct {
	feud.BaseBarony
}

var BAughagower阿哈乔韦尔 feud.Barony = &阿哈乔韦尔AughagowerBarony{}

func init() {
    f := BAughagower阿哈乔韦尔.(*阿哈乔韦尔AughagowerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aughagower",
		TitleName: "阿哈乔韦尔",
		TitleCode: "b_aughagower",
	}
}
