package trans-portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨莫杰德SamodedBarony struct {
	feud.BaseBarony
}

var BSamoded萨莫杰德 feud.Barony = &萨莫杰德SamodedBarony{}

func init() {
    f := BSamoded萨莫杰德.(*萨莫杰德SamodedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samoded",
		TitleName: "萨莫杰德",
		TitleCode: "b_samoded",
	}
}
