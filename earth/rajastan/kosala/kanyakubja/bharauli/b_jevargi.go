package bharauli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祗伐耆JevargiBarony struct {
	feud.BaseBarony
}

var BJevargi祗伐耆 feud.Barony = &祗伐耆JevargiBarony{}

func init() {
    f := BJevargi祗伐耆.(*祗伐耆JevargiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jevargi",
		TitleName: "祗伐耆",
		TitleCode: "b_jevargi",
	}
}
