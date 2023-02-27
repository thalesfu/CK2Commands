package neuchatel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣伊米耶StimierBarony struct {
	feud.BaseBarony
}

var BStimier圣伊米耶 feud.Barony = &圣伊米耶StimierBarony{}

func init() {
    f := BStimier圣伊米耶.(*圣伊米耶StimierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stimier",
		TitleName: "圣伊米耶",
		TitleCode: "b_stimier",
	}
}
