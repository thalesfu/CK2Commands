package vanema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁梅内RumenBarony struct {
	feud.BaseBarony
}

var BRumen鲁梅内 feud.Barony = &鲁梅内RumenBarony{}

func init() {
    f := BRumen鲁梅内.(*鲁梅内RumenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rumen",
		TitleName: "鲁梅内",
		TitleCode: "b_rumen",
	}
}
