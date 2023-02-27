package more

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥勒姆AlemBarony struct {
	feud.BaseBarony
}

var BAlem奥勒姆 feud.Barony = &奥勒姆AlemBarony{}

func init() {
    f := BAlem奥勒姆.(*奥勒姆AlemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alem",
		TitleName: "奥勒姆",
		TitleCode: "b_alem",
	}
}
