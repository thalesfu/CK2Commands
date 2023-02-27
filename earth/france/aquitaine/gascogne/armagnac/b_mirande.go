package armagnac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米朗德MirandeBarony struct {
	feud.BaseBarony
}

var BMirande米朗德 feud.Barony = &米朗德MirandeBarony{}

func init() {
    f := BMirande米朗德.(*米朗德MirandeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirande",
		TitleName: "米朗德",
		TitleCode: "b_mirande",
	}
}
