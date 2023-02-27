package multan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茂罗三部卢MulasthanaBarony struct {
	feud.BaseBarony
}

var BMulasthana茂罗三部卢 feud.Barony = &茂罗三部卢MulasthanaBarony{}

func init() {
    f := BMulasthana茂罗三部卢.(*茂罗三部卢MulasthanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mulasthana",
		TitleName: "茂罗三部卢",
		TitleCode: "b_mulasthana",
	}
}
