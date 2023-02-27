package montpellier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马格洛讷MagueloneBarony struct {
	feud.BaseBarony
}

var BMaguelone马格洛讷 feud.Barony = &马格洛讷MagueloneBarony{}

func init() {
    f := BMaguelone马格洛讷.(*马格洛讷MagueloneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maguelone",
		TitleName: "马格洛讷",
		TitleCode: "b_maguelone",
	}
}
