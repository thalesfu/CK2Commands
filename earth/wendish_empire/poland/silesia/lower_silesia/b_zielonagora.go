package lower_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绿山城ZielonagoraBarony struct {
	feud.BaseBarony
}

var BZielonagora绿山城 feud.Barony = &绿山城ZielonagoraBarony{}

func init() {
    f := BZielonagora绿山城.(*绿山城ZielonagoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zielonagora",
		TitleName: "绿山城",
		TitleCode: "b_zielonagora",
	}
}
