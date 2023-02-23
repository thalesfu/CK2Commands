package mali

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布古尼BougouniBarony struct {
	feud.BaseBarony
}

var BBougouni布古尼 feud.Barony = &布古尼BougouniBarony{}

func init() {
	f := BBougouni布古尼.(*布古尼BougouniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bougouni",
		TitleName: "布古尼",
		TitleCode: "b_bougouni",
	}
}
