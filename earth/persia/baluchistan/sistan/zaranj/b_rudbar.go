package zaranj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁德巴尔RudbarBarony struct {
	feud.BaseBarony
}

var BRudbar鲁德巴尔 feud.Barony = &鲁德巴尔RudbarBarony{}

func init() {
    f := BRudbar鲁德巴尔.(*鲁德巴尔RudbarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rudbar",
		TitleName: "鲁德巴尔",
		TitleCode: "b_rudbar",
	}
}
