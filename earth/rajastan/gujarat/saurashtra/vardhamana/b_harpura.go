package vardhamana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曷罗补罗HarpuraBarony struct {
	feud.BaseBarony
}

var BHarpura曷罗补罗 feud.Barony = &曷罗补罗HarpuraBarony{}

func init() {
    f := BHarpura曷罗补罗.(*曷罗补罗HarpuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harpura",
		TitleName: "曷罗补罗",
		TitleCode: "b_harpura",
	}
}
