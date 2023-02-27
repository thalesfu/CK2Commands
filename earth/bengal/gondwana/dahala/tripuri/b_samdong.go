package tripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 僧东SamdongBarony struct {
	feud.BaseBarony
}

var BSamdong僧东 feud.Barony = &僧东SamdongBarony{}

func init() {
    f := BSamdong僧东.(*僧东SamdongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samdong",
		TitleName: "僧东",
		TitleCode: "b_samdong",
	}
}
