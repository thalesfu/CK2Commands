package bremen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布卢门塔尔BlumenthalBarony struct {
	feud.BaseBarony
}

var BBlumenthal布卢门塔尔 feud.Barony = &布卢门塔尔BlumenthalBarony{}

func init() {
	f := BBlumenthal布卢门塔尔.(*布卢门塔尔BlumenthalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "blumenthal",
		TitleName: "布卢门塔尔",
		TitleCode: "b_blumenthal",
	}
}
