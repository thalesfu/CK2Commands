package vasterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮特奥PiteaBarony struct {
	feud.BaseBarony
}

var BPitea皮特奥 feud.Barony = &皮特奥PiteaBarony{}

func init() {
    f := BPitea皮特奥.(*皮特奥PiteaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pitea",
		TitleName: "皮特奥",
		TitleCode: "b_pitea",
	}
}
