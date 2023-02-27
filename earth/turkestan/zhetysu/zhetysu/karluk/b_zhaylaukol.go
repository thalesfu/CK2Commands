package karluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎伊劳科利ZhaylaukolBarony struct {
	feud.BaseBarony
}

var BZhaylaukol扎伊劳科利 feud.Barony = &扎伊劳科利ZhaylaukolBarony{}

func init() {
    f := BZhaylaukol扎伊劳科利.(*扎伊劳科利ZhaylaukolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhaylaukol",
		TitleName: "扎伊劳科利",
		TitleCode: "b_zhaylaukol",
	}
}
