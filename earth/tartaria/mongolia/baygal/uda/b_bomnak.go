package uda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博姆纳克BomnakBarony struct {
	feud.BaseBarony
}

var BBomnak博姆纳克 feud.Barony = &博姆纳克BomnakBarony{}

func init() {
    f := BBomnak博姆纳克.(*博姆纳克BomnakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bomnak",
		TitleName: "博姆纳克",
		TitleCode: "b_bomnak",
	}
}
