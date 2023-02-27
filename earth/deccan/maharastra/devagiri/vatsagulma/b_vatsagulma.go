package vatsagulma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆蹉崛摩VatsagulmaBarony struct {
	feud.BaseBarony
}

var BVatsagulma婆蹉崛摩 feud.Barony = &婆蹉崛摩VatsagulmaBarony{}

func init() {
    f := BVatsagulma婆蹉崛摩.(*婆蹉崛摩VatsagulmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vatsagulma",
		TitleName: "婆蹉崛摩",
		TitleCode: "b_vatsagulma",
	}
}
