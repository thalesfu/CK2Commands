package ross

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿普尔克罗斯ApplecrossBarony struct {
	feud.BaseBarony
}

var BApplecross阿普尔克罗斯 feud.Barony = &阿普尔克罗斯ApplecrossBarony{}

func init() {
    f := BApplecross阿普尔克罗斯.(*阿普尔克罗斯ApplecrossBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "applecross",
		TitleName: "阿普尔克罗斯",
		TitleCode: "b_applecross",
	}
}
