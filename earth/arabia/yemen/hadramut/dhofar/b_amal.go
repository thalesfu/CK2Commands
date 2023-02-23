package dhofar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿迈勒AmalBarony struct {
	feud.BaseBarony
}

var BAmal阿迈勒 feud.Barony = &阿迈勒AmalBarony{}

func init() {
	f := BAmal阿迈勒.(*阿迈勒AmalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amal",
		TitleName: "阿迈勒",
		TitleCode: "b_amal",
	}
}
