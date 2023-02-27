package sieradzka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗兹普沙RozprzaBarony struct {
	feud.BaseBarony
}

var BRozprza罗兹普沙 feud.Barony = &罗兹普沙RozprzaBarony{}

func init() {
    f := BRozprza罗兹普沙.(*罗兹普沙RozprzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rozprza",
		TitleName: "罗兹普沙",
		TitleCode: "b_rozprza",
	}
}
