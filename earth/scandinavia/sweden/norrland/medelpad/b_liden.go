package medelpad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利登LidenBarony struct {
	feud.BaseBarony
}

var BLiden利登 feud.Barony = &利登LidenBarony{}

func init() {
	f := BLiden利登.(*利登LidenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liden",
		TitleName: "利登",
		TitleCode: "b_liden",
	}
}
