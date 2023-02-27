package kafirkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比诺特BilotBarony struct {
	feud.BaseBarony
}

var BBilot比诺特 feud.Barony = &比诺特BilotBarony{}

func init() {
    f := BBilot比诺特.(*比诺特BilotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilot",
		TitleName: "比诺特",
		TitleCode: "b_bilot",
	}
}
