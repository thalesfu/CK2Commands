package vaud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌希OuchyBarony struct {
	feud.BaseBarony
}

var BOuchy乌希 feud.Barony = &乌希OuchyBarony{}

func init() {
    f := BOuchy乌希.(*乌希OuchyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouchy",
		TitleName: "乌希",
		TitleCode: "b_ouchy",
	}
}
