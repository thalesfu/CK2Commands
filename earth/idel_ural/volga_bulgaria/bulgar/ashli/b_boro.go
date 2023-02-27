package ashli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯勒BoroBarony struct {
	feud.BaseBarony
}

var BBoro伯勒 feud.Barony = &伯勒BoroBarony{}

func init() {
    f := BBoro伯勒.(*伯勒BoroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boro",
		TitleName: "伯勒",
		TitleCode: "b_boro",
	}
}
