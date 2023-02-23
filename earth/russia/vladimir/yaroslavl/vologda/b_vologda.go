package vologda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃洛格达VologdaBarony struct {
	feud.BaseBarony
}

var BVologda沃洛格达 feud.Barony = &沃洛格达VologdaBarony{}

func init() {
	f := BVologda沃洛格达.(*沃洛格达VologdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vologda",
		TitleName: "沃洛格达",
		TitleCode: "b_vologda",
	}
}
