package sunnmore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔贡BorgundBarony struct {
	feud.BaseBarony
}

var BBorgund博尔贡 feud.Barony = &博尔贡BorgundBarony{}

func init() {
    f := BBorgund博尔贡.(*博尔贡BorgundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borgund",
		TitleName: "博尔贡",
		TitleCode: "b_borgund",
	}
}
