package selija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥齐凯尔OzikelBarony struct {
	feud.BaseBarony
}

var BOzikel奥齐凯尔 feud.Barony = &奥齐凯尔OzikelBarony{}

func init() {
    f := BOzikel奥齐凯尔.(*奥齐凯尔OzikelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ozikel",
		TitleName: "奥齐凯尔",
		TitleCode: "b_ozikel",
	}
}
