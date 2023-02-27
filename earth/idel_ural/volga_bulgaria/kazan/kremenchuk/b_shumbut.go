package kremenchuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 顺布特ShumbutBarony struct {
	feud.BaseBarony
}

var BShumbut顺布特 feud.Barony = &顺布特ShumbutBarony{}

func init() {
    f := BShumbut顺布特.(*顺布特ShumbutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shumbut",
		TitleName: "顺布特",
		TitleCode: "b_shumbut",
	}
}
