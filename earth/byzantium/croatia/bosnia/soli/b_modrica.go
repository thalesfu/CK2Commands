package soli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫德里查ModricaBarony struct {
	feud.BaseBarony
}

var BModrica莫德里查 feud.Barony = &莫德里查ModricaBarony{}

func init() {
	f := BModrica莫德里查.(*莫德里查ModricaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "modrica",
		TitleName: "莫德里查",
		TitleCode: "b_modrica",
	}
}
