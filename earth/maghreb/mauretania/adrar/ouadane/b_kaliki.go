package ouadane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利基KalikiBarony struct {
	feud.BaseBarony
}

var BKaliki卡利基 feud.Barony = &卡利基KalikiBarony{}

func init() {
    f := BKaliki卡利基.(*卡利基KalikiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaliki",
		TitleName: "卡利基",
		TitleCode: "b_kaliki",
	}
}
