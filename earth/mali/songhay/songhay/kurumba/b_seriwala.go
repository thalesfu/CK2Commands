package kurumba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞里瓦拉SeriwalaBarony struct {
	feud.BaseBarony
}

var BSeriwala塞里瓦拉 feud.Barony = &塞里瓦拉SeriwalaBarony{}

func init() {
	f := BSeriwala塞里瓦拉.(*塞里瓦拉SeriwalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seriwala",
		TitleName: "塞里瓦拉",
		TitleCode: "b_seriwala",
	}
}
