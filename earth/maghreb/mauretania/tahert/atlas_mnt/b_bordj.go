package atlas_mnt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔吉BordjBarony struct {
	feud.BaseBarony
}

var BBordj布尔吉 feud.Barony = &布尔吉BordjBarony{}

func init() {
    f := BBordj布尔吉.(*布尔吉BordjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bordj",
		TitleName: "布尔吉",
		TitleCode: "b_bordj",
	}
}
