package vychegda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维切格达VychegdaBarony struct {
	feud.BaseBarony
}

var BVychegda维切格达 feud.Barony = &维切格达VychegdaBarony{}

func init() {
    f := BVychegda维切格达.(*维切格达VychegdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vychegda",
		TitleName: "维切格达",
		TitleCode: "b_vychegda",
	}
}
