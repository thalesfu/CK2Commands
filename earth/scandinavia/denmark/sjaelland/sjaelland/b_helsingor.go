package sjaelland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫尔辛格HelsingorBarony struct {
	feud.BaseBarony
}

var BHelsingor赫尔辛格 feud.Barony = &赫尔辛格HelsingorBarony{}

func init() {
    f := BHelsingor赫尔辛格.(*赫尔辛格HelsingorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "helsingor",
		TitleName: "赫尔辛格",
		TitleCode: "b_helsingor",
	}
}
