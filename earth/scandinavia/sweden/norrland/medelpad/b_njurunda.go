package medelpad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纽伦达NjurundaBarony struct {
	feud.BaseBarony
}

var BNjurunda纽伦达 feud.Barony = &纽伦达NjurundaBarony{}

func init() {
    f := BNjurunda纽伦达.(*纽伦达NjurundaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "njurunda",
		TitleName: "纽伦达",
		TitleCode: "b_njurunda",
	}
}
