package ascalon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布兰彻贾德BlanchegardeBarony struct {
	feud.BaseBarony
}

var BBlanchegarde布兰彻贾德 feud.Barony = &布兰彻贾德BlanchegardeBarony{}

func init() {
    f := BBlanchegarde布兰彻贾德.(*布兰彻贾德BlanchegardeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "blanchegarde",
		TitleName: "布兰彻贾德",
		TitleCode: "b_blanchegarde",
	}
}
