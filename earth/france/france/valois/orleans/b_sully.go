package orleans

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叙利SullyBarony struct {
	feud.BaseBarony
}

var BSully叙利 feud.Barony = &叙利SullyBarony{}

func init() {
    f := BSully叙利.(*叙利SullyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sully",
		TitleName: "叙利",
		TitleCode: "b_sully",
	}
}
