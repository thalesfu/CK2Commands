package bilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古阿姆GouamBarony struct {
	feud.BaseBarony
}

var BGouam古阿姆 feud.Barony = &古阿姆GouamBarony{}

func init() {
    f := BGouam古阿姆.(*古阿姆GouamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gouam",
		TitleName: "古阿姆",
		TitleCode: "b_gouam",
	}
}
