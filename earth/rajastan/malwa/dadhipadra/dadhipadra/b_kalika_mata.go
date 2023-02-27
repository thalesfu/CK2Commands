package dadhipadra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽利伽摩多Kalika_mataBarony struct {
	feud.BaseBarony
}

var BKalika_mata伽利伽摩多 feud.Barony = &伽利伽摩多Kalika_mataBarony{}

func init() {
    f := BKalika_mata伽利伽摩多.(*伽利伽摩多Kalika_mataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalika_mata",
		TitleName: "伽利伽摩多",
		TitleCode: "b_kalika_mata",
	}
}
