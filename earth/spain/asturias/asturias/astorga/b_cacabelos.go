package astorga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡卡韦洛斯CacabelosBarony struct {
	feud.BaseBarony
}

var BCacabelos卡卡韦洛斯 feud.Barony = &卡卡韦洛斯CacabelosBarony{}

func init() {
    f := BCacabelos卡卡韦洛斯.(*卡卡韦洛斯CacabelosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cacabelos",
		TitleName: "卡卡韦洛斯",
		TitleCode: "b_cacabelos",
	}
}
