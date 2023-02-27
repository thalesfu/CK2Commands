package ugra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 武克特尔VuktylBarony struct {
	feud.BaseBarony
}

var BVuktyl武克特尔 feud.Barony = &武克特尔VuktylBarony{}

func init() {
    f := BVuktyl武克特尔.(*武克特尔VuktylBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vuktyl",
		TitleName: "武克特尔",
		TitleCode: "b_vuktyl",
	}
}
