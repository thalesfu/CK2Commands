package brie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂耶里堡ChateauthierryBarony struct {
	feud.BaseBarony
}

var BChateauthierry蒂耶里堡 feud.Barony = &蒂耶里堡ChateauthierryBarony{}

func init() {
	f := BChateauthierry蒂耶里堡.(*蒂耶里堡ChateauthierryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chateauthierry",
		TitleName: "蒂耶里堡",
		TitleCode: "b_chateauthierry",
	}
}
