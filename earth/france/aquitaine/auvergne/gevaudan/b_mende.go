package gevaudan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 芒德MendeBarony struct {
	feud.BaseBarony
}

var BMende芒德 feud.Barony = &芒德MendeBarony{}

func init() {
    f := BMende芒德.(*芒德MendeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mende",
		TitleName: "芒德",
		TitleCode: "b_mende",
	}
}
