package anhalt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科斯维希CoswigBarony struct {
	feud.BaseBarony
}

var BCoswig科斯维希 feud.Barony = &科斯维希CoswigBarony{}

func init() {
    f := BCoswig科斯维希.(*科斯维希CoswigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coswig",
		TitleName: "科斯维希",
		TitleCode: "b_coswig",
	}
}
