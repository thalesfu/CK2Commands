package syrte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西瑙安SinauenBarony struct {
	feud.BaseBarony
}

var BSinauen西瑙安 feud.Barony = &西瑙安SinauenBarony{}

func init() {
    f := BSinauen西瑙安.(*西瑙安SinauenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sinauen",
		TitleName: "西瑙安",
		TitleCode: "b_sinauen",
	}
}
