package kuopio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉韦西KallavesiBarony struct {
	feud.BaseBarony
}

var BKallavesi卡拉韦西 feud.Barony = &卡拉韦西KallavesiBarony{}

func init() {
	f := BKallavesi卡拉韦西.(*卡拉韦西KallavesiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kallavesi",
		TitleName: "卡拉韦西",
		TitleCode: "b_kallavesi",
	}
}
