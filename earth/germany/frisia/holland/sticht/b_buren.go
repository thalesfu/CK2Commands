package sticht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比伦BurenBarony struct {
	feud.BaseBarony
}

var BBuren比伦 feud.Barony = &比伦BurenBarony{}

func init() {
	f := BBuren比伦.(*比伦BurenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buren",
		TitleName: "比伦",
		TitleCode: "b_buren",
	}
}
