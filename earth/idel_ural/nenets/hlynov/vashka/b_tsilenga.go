package vashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐连加TsilengaBarony struct {
	feud.BaseBarony
}

var BTsilenga齐连加 feud.Barony = &齐连加TsilengaBarony{}

func init() {
    f := BTsilenga齐连加.(*齐连加TsilengaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsilenga",
		TitleName: "齐连加",
		TitleCode: "b_tsilenga",
	}
}
