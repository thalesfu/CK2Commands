package bobo_dyulasso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本通巴BountoumbaBarony struct {
	feud.BaseBarony
}

var BBountoumba本通巴 feud.Barony = &本通巴BountoumbaBarony{}

func init() {
    f := BBountoumba本通巴.(*本通巴BountoumbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bountoumba",
		TitleName: "本通巴",
		TitleCode: "b_bountoumba",
	}
}
