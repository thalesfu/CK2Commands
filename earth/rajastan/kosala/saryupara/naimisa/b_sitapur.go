package naimisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 私多城SitapurBarony struct {
	feud.BaseBarony
}

var BSitapur私多城 feud.Barony = &私多城SitapurBarony{}

func init() {
    f := BSitapur私多城.(*私多城SitapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sitapur",
		TitleName: "私多城",
		TitleCode: "b_sitapur",
	}
}
