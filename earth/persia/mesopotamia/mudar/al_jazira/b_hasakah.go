package al_jazira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈塞克HasakahBarony struct {
	feud.BaseBarony
}

var BHasakah哈塞克 feud.Barony = &哈塞克HasakahBarony{}

func init() {
    f := BHasakah哈塞克.(*哈塞克HasakahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hasakah",
		TitleName: "哈塞克",
		TitleCode: "b_hasakah",
	}
}
