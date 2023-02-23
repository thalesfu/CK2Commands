package glamorgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马格姆MargamBarony struct {
	feud.BaseBarony
}

var BMargam马格姆 feud.Barony = &马格姆MargamBarony{}

func init() {
	f := BMargam马格姆.(*马格姆MargamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "margam",
		TitleName: "马格姆",
		TitleCode: "b_margam",
	}
}
