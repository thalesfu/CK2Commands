package aukshayts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅迪宁凯MedininkaiBarony struct {
	feud.BaseBarony
}

var BMedininkai梅迪宁凯 feud.Barony = &梅迪宁凯MedininkaiBarony{}

func init() {
    f := BMedininkai梅迪宁凯.(*梅迪宁凯MedininkaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medininkai",
		TitleName: "梅迪宁凯",
		TitleCode: "b_medininkai",
	}
}
