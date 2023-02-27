package lhoyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珞隅LhoyuBarony struct {
	feud.BaseBarony
}

var BLhoyu珞隅 feud.Barony = &珞隅LhoyuBarony{}

func init() {
    f := BLhoyu珞隅.(*珞隅LhoyuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lhoyu",
		TitleName: "珞隅",
		TitleCode: "b_lhoyu",
	}
}
