package mosul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴克迪达BakhdidaBarony struct {
	feud.BaseBarony
}

var BBakhdida巴克迪达 feud.Barony = &巴克迪达BakhdidaBarony{}

func init() {
    f := BBakhdida巴克迪达.(*巴克迪达BakhdidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakhdida",
		TitleName: "巴克迪达",
		TitleCode: "b_bakhdida",
	}
}
