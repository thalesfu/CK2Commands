package vorotynsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥多耶夫OdoyevBarony struct {
	feud.BaseBarony
}

var BOdoyev奥多耶夫 feud.Barony = &奥多耶夫OdoyevBarony{}

func init() {
    f := BOdoyev奥多耶夫.(*奥多耶夫OdoyevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "odoyev",
		TitleName: "奥多耶夫",
		TitleCode: "b_odoyev",
	}
}
