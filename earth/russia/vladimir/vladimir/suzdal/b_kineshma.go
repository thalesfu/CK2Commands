package suzdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基涅什马KineshmaBarony struct {
	feud.BaseBarony
}

var BKineshma基涅什马 feud.Barony = &基涅什马KineshmaBarony{}

func init() {
	f := BKineshma基涅什马.(*基涅什马KineshmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kineshma",
		TitleName: "基涅什马",
		TitleCode: "b_kineshma",
	}
}
