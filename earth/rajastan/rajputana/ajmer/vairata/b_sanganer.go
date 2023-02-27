package vairata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑格内尔SanganerBarony struct {
	feud.BaseBarony
}

var BSanganer桑格内尔 feud.Barony = &桑格内尔SanganerBarony{}

func init() {
    f := BSanganer桑格内尔.(*桑格内尔SanganerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanganer",
		TitleName: "桑格内尔",
		TitleCode: "b_sanganer",
	}
}
