package vannes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛克米内LocmineBarony struct {
	feud.BaseBarony
}

var BLocmine洛克米内 feud.Barony = &洛克米内LocmineBarony{}

func init() {
	f := BLocmine洛克米内.(*洛克米内LocmineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "locmine",
		TitleName: "洛克米内",
		TitleCode: "b_locmine",
	}
}
