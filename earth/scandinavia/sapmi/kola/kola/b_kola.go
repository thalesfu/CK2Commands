package kola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科拉KolaBarony struct {
	feud.BaseBarony
}

var BKola科拉 feud.Barony = &科拉KolaBarony{}

func init() {
    f := BKola科拉.(*科拉KolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kola",
		TitleName: "科拉",
		TitleCode: "b_kola",
	}
}
