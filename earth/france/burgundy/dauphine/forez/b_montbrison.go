package forez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙布里松MontbrisonBarony struct {
	feud.BaseBarony
}

var BMontbrison蒙布里松 feud.Barony = &蒙布里松MontbrisonBarony{}

func init() {
    f := BMontbrison蒙布里松.(*蒙布里松MontbrisonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montbrison",
		TitleName: "蒙布里松",
		TitleCode: "b_montbrison",
	}
}
