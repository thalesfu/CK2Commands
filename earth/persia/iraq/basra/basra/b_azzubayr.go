package basra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖拜尔AzzubayrBarony struct {
	feud.BaseBarony
}

var BAzzubayr祖拜尔 feud.Barony = &祖拜尔AzzubayrBarony{}

func init() {
    f := BAzzubayr祖拜尔.(*祖拜尔AzzubayrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azzubayr",
		TitleName: "祖拜尔",
		TitleCode: "b_azzubayr",
	}
}
