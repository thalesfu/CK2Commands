package vukovar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基巴莱CibalaeBarony struct {
	feud.BaseBarony
}

var BCibalae基巴莱 feud.Barony = &基巴莱CibalaeBarony{}

func init() {
	f := BCibalae基巴莱.(*基巴莱CibalaeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cibalae",
		TitleName: "基巴莱",
		TitleCode: "b_cibalae",
	}
}
