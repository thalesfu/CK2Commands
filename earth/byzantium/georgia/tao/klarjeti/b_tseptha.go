package klarjeti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 采普萨TsepthaBarony struct {
	feud.BaseBarony
}

var BTseptha采普萨 feud.Barony = &采普萨TsepthaBarony{}

func init() {
	f := BTseptha采普萨.(*采普萨TsepthaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tseptha",
		TitleName: "采普萨",
		TitleCode: "b_tseptha",
	}
}
