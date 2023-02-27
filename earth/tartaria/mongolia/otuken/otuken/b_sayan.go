package otuken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨彦SayanBarony struct {
	feud.BaseBarony
}

var BSayan萨彦 feud.Barony = &萨彦SayanBarony{}

func init() {
    f := BSayan萨彦.(*萨彦SayanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sayan",
		TitleName: "萨彦",
		TitleCode: "b_sayan",
	}
}
