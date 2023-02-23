package zeila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉吉加JijigaBarony struct {
	feud.BaseBarony
}

var BJijiga吉吉加 feud.Barony = &吉吉加JijigaBarony{}

func init() {
	f := BJijiga吉吉加.(*吉吉加JijigaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jijiga",
		TitleName: "吉吉加",
		TitleCode: "b_jijiga",
	}
}
