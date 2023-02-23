package galloway

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 威格敦WigtownBarony struct {
	feud.BaseBarony
}

var BWigtown威格敦 feud.Barony = &威格敦WigtownBarony{}

func init() {
	f := BWigtown威格敦.(*威格敦WigtownBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wigtown",
		TitleName: "威格敦",
		TitleCode: "b_wigtown",
	}
}
