package perugia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔夏诺MarscianoBarony struct {
	feud.BaseBarony
}

var BMarsciano马尔夏诺 feud.Barony = &马尔夏诺MarscianoBarony{}

func init() {
	f := BMarsciano马尔夏诺.(*马尔夏诺MarscianoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marsciano",
		TitleName: "马尔夏诺",
		TitleCode: "b_marsciano",
	}
}
