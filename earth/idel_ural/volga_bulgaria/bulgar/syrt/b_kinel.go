package syrt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基涅利KinelBarony struct {
	feud.BaseBarony
}

var BKinel基涅利 feud.Barony = &基涅利KinelBarony{}

func init() {
    f := BKinel基涅利.(*基涅利KinelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kinel",
		TitleName: "基涅利",
		TitleCode: "b_kinel",
	}
}
