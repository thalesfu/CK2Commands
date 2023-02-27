package suvarnagram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩迦伊城BokainagarBarony struct {
	feud.BaseBarony
}

var BBokainagar菩迦伊城 feud.Barony = &菩迦伊城BokainagarBarony{}

func init() {
    f := BBokainagar菩迦伊城.(*菩迦伊城BokainagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bokainagar",
		TitleName: "菩迦伊城",
		TitleCode: "b_bokainagar",
	}
}
