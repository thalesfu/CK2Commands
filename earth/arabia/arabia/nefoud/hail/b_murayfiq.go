package hail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆赖菲格MurayfiqBarony struct {
	feud.BaseBarony
}

var BMurayfiq穆赖菲格 feud.Barony = &穆赖菲格MurayfiqBarony{}

func init() {
    f := BMurayfiq穆赖菲格.(*穆赖菲格MurayfiqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murayfiq",
		TitleName: "穆赖菲格",
		TitleCode: "b_murayfiq",
	}
}
