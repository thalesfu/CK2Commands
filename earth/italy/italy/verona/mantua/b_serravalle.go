package mantua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞拉瓦莱SerravalleBarony struct {
	feud.BaseBarony
}

var BSerravalle塞拉瓦莱 feud.Barony = &塞拉瓦莱SerravalleBarony{}

func init() {
    f := BSerravalle塞拉瓦莱.(*塞拉瓦莱SerravalleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "serravalle",
		TitleName: "塞拉瓦莱",
		TitleCode: "b_serravalle",
	}
}
