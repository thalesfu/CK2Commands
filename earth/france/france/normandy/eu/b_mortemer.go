package eu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫特梅MortemerBarony struct {
	feud.BaseBarony
}

var BMortemer莫特梅 feud.Barony = &莫特梅MortemerBarony{}

func init() {
    f := BMortemer莫特梅.(*莫特梅MortemerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mortemer",
		TitleName: "莫特梅",
		TitleCode: "b_mortemer",
	}
}
