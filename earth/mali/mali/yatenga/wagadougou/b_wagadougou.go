package wagadougou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦加杜古WagadougouBarony struct {
	feud.BaseBarony
}

var BWagadougou瓦加杜古 feud.Barony = &瓦加杜古WagadougouBarony{}

func init() {
	f := BWagadougou瓦加杜古.(*瓦加杜古WagadougouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wagadougou",
		TitleName: "瓦加杜古",
		TitleCode: "b_wagadougou",
	}
}
