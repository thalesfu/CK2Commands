package nyingchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 百巴BebaBarony struct {
	feud.BaseBarony
}

var BBeba百巴 feud.Barony = &百巴BebaBarony{}

func init() {
    f := BBeba百巴.(*百巴BebaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beba",
		TitleName: "百巴",
		TitleCode: "b_beba",
	}
}
