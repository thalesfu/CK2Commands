package nandapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祇耶城JeyporeBarony struct {
	feud.BaseBarony
}

var BJeypore祇耶城 feud.Barony = &祇耶城JeyporeBarony{}

func init() {
    f := BJeypore祇耶城.(*祇耶城JeyporeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jeypore",
		TitleName: "祇耶城",
		TitleCode: "b_jeypore",
	}
}
