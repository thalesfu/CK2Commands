package egiin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 额金EgiinBarony struct {
	feud.BaseBarony
}

var BEgiin额金 feud.Barony = &额金EgiinBarony{}

func init() {
	f := BEgiin额金.(*额金EgiinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egiin",
		TitleName: "额金",
		TitleCode: "b_egiin",
	}
}
