package maldives

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼都MundooBarony struct {
	feud.BaseBarony
}

var BMundoo曼都 feud.Barony = &曼都MundooBarony{}

func init() {
    f := BMundoo曼都.(*曼都MundooBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mundoo",
		TitleName: "曼都",
		TitleCode: "b_mundoo",
	}
}
