package west_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉斯拉瓦KraslavaBarony struct {
	feud.BaseBarony
}

var BKraslava克拉斯拉瓦 feud.Barony = &克拉斯拉瓦KraslavaBarony{}

func init() {
    f := BKraslava克拉斯拉瓦.(*克拉斯拉瓦KraslavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kraslava",
		TitleName: "克拉斯拉瓦",
		TitleCode: "b_kraslava",
	}
}
