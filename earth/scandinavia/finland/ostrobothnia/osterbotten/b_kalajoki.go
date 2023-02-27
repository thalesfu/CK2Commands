package osterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉约基KalajokiBarony struct {
	feud.BaseBarony
}

var BKalajoki卡拉约基 feud.Barony = &卡拉约基KalajokiBarony{}

func init() {
    f := BKalajoki卡拉约基.(*卡拉约基KalajokiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalajoki",
		TitleName: "卡拉约基",
		TitleCode: "b_kalajoki",
	}
}
