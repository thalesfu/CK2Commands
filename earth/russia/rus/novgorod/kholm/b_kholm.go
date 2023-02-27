package kholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔姆KholmBarony struct {
	feud.BaseBarony
}

var BKholm霍尔姆 feud.Barony = &霍尔姆KholmBarony{}

func init() {
    f := BKholm霍尔姆.(*霍尔姆KholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kholm",
		TitleName: "霍尔姆",
		TitleCode: "b_kholm",
	}
}
