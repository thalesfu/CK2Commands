package koshma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科沙姆KoshmaBarony struct {
	feud.BaseBarony
}

var BKoshma科沙姆 feud.Barony = &科沙姆KoshmaBarony{}

func init() {
    f := BKoshma科沙姆.(*科沙姆KoshmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koshma",
		TitleName: "科沙姆",
		TitleCode: "b_koshma",
	}
}
