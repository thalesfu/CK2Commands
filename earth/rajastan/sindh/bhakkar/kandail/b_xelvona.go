package kandail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切符那XelvonaBarony struct {
	feud.BaseBarony
}

var BXelvona切符那 feud.Barony = &切符那XelvonaBarony{}

func init() {
	f := BXelvona切符那.(*切符那XelvonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xelvona",
		TitleName: "切符那",
		TitleCode: "b_xelvona",
	}
}
