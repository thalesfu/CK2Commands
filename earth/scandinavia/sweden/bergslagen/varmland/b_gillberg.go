package varmland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔贝里GillbergBarony struct {
	feud.BaseBarony
}

var BGillberg伊尔贝里 feud.Barony = &伊尔贝里GillbergBarony{}

func init() {
    f := BGillberg伊尔贝里.(*伊尔贝里GillbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gillberg",
		TitleName: "伊尔贝里",
		TitleCode: "b_gillberg",
	}
}
