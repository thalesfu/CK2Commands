package monemvasia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉雄GythioBarony struct {
	feud.BaseBarony
}

var BGythio吉雄 feud.Barony = &吉雄GythioBarony{}

func init() {
    f := BGythio吉雄.(*吉雄GythioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gythio",
		TitleName: "吉雄",
		TitleCode: "b_gythio",
	}
}
