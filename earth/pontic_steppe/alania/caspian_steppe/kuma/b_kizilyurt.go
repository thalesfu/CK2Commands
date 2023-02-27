package kuma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒尤尔特KizilyurtBarony struct {
	feud.BaseBarony
}

var BKizilyurt克孜勒尤尔特 feud.Barony = &克孜勒尤尔特KizilyurtBarony{}

func init() {
    f := BKizilyurt克孜勒尤尔特.(*克孜勒尤尔特KizilyurtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kizilyurt",
		TitleName: "克孜勒尤尔特",
		TitleCode: "b_kizilyurt",
	}
}
