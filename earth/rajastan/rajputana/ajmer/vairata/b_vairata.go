package vairata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠罗多VairataBarony struct {
	feud.BaseBarony
}

var BVairata吠罗多 feud.Barony = &吠罗多VairataBarony{}

func init() {
    f := BVairata吠罗多.(*吠罗多VairataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vairata",
		TitleName: "吠罗多",
		TitleCode: "b_vairata",
	}
}
