package tortosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马加特MargatBarony struct {
	feud.BaseBarony
}

var BMargat马加特 feud.Barony = &马加特MargatBarony{}

func init() {
    f := BMargat马加特.(*马加特MargatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "margat",
		TitleName: "马加特",
		TitleCode: "b_margat",
	}
}
