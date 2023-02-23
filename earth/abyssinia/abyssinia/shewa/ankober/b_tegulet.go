package ankober

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特古勒TeguletBarony struct {
	feud.BaseBarony
}

var BTegulet特古勒 feud.Barony = &特古勒TeguletBarony{}

func init() {
	f := BTegulet特古勒.(*特古勒TeguletBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tegulet",
		TitleName: "特古勒",
		TitleCode: "b_tegulet",
	}
}
