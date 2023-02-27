package kondana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 底拘那TikonaBarony struct {
	feud.BaseBarony
}

var BTikona底拘那 feud.Barony = &底拘那TikonaBarony{}

func init() {
    f := BTikona底拘那.(*底拘那TikonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tikona",
		TitleName: "底拘那",
		TitleCode: "b_tikona",
	}
}
