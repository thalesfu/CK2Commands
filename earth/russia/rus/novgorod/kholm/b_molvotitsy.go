package kholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔沃季齐MolvotitsyBarony struct {
	feud.BaseBarony
}

var BMolvotitsy莫尔沃季齐 feud.Barony = &莫尔沃季齐MolvotitsyBarony{}

func init() {
    f := BMolvotitsy莫尔沃季齐.(*莫尔沃季齐MolvotitsyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "molvotitsy",
		TitleName: "莫尔沃季齐",
		TitleCode: "b_molvotitsy",
	}
}
