package bome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波密BomeBarony struct {
	feud.BaseBarony
}

var BBome波密 feud.Barony = &波密BomeBarony{}

func init() {
    f := BBome波密.(*波密BomeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bome",
		TitleName: "波密",
		TitleCode: "b_bome",
	}
}
