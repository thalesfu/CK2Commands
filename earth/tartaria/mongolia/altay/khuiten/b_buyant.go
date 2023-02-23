package khuiten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布彦特BuyantBarony struct {
	feud.BaseBarony
}

var BBuyant布彦特 feud.Barony = &布彦特BuyantBarony{}

func init() {
	f := BBuyant布彦特.(*布彦特BuyantBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buyant",
		TitleName: "布彦特",
		TitleCode: "b_buyant",
	}
}
