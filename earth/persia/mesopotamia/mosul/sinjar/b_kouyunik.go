package sinjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阔玉尼克KouyunikBarony struct {
	feud.BaseBarony
}

var BKouyunik阔玉尼克 feud.Barony = &阔玉尼克KouyunikBarony{}

func init() {
	f := BKouyunik阔玉尼克.(*阔玉尼克KouyunikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kouyunik",
		TitleName: "阔玉尼克",
		TitleCode: "b_kouyunik",
	}
}
