package karvuna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦内茨VenetsBarony struct {
	feud.BaseBarony
}

var BVenets韦内茨 feud.Barony = &韦内茨VenetsBarony{}

func init() {
    f := BVenets韦内茨.(*韦内茨VenetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "venets",
		TitleName: "韦内茨",
		TitleCode: "b_venets",
	}
}
