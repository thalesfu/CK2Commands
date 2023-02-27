package aleppo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿勒颇AleppoBarony struct {
	feud.BaseBarony
}

var BAleppo阿勒颇 feud.Barony = &阿勒颇AleppoBarony{}

func init() {
    f := BAleppo阿勒颇.(*阿勒颇AleppoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aleppo",
		TitleName: "阿勒颇",
		TitleCode: "b_aleppo",
	}
}
