package trier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施蓬海姆SponheimBarony struct {
	feud.BaseBarony
}

var BSponheim施蓬海姆 feud.Barony = &施蓬海姆SponheimBarony{}

func init() {
    f := BSponheim施蓬海姆.(*施蓬海姆SponheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sponheim",
		TitleName: "施蓬海姆",
		TitleCode: "b_sponheim",
	}
}
