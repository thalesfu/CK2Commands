package fejer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔MorBarony struct {
	feud.BaseBarony
}

var BMor莫尔 feud.Barony = &莫尔MorBarony{}

func init() {
    f := BMor莫尔.(*莫尔MorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mor",
		TitleName: "莫尔",
		TitleCode: "b_mor",
	}
}
