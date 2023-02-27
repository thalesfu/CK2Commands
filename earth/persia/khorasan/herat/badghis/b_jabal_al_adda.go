package badghis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾巴勒阿达Jabal_al_addaBarony struct {
	feud.BaseBarony
}

var BJabal_al_adda贾巴勒阿达 feud.Barony = &贾巴勒阿达Jabal_al_addaBarony{}

func init() {
    f := BJabal_al_adda贾巴勒阿达.(*贾巴勒阿达Jabal_al_addaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jabal_al_adda",
		TitleName: "贾巴勒阿达",
		TitleCode: "b_jabal_al_adda",
	}
}
