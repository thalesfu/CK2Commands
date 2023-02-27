package aylik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙湾ShawanBarony struct {
	feud.BaseBarony
}

var BShawan沙湾 feud.Barony = &沙湾ShawanBarony{}

func init() {
    f := BShawan沙湾.(*沙湾ShawanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shawan",
		TitleName: "沙湾",
		TitleCode: "b_shawan",
	}
}
