package heves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫维什HevesBarony struct {
	feud.BaseBarony
}

var BHeves赫维什 feud.Barony = &赫维什HevesBarony{}

func init() {
	f := BHeves赫维什.(*赫维什HevesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heves",
		TitleName: "赫维什",
		TitleCode: "b_heves",
	}
}
