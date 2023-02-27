package lukomorie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利米乌斯KalmiusBarony struct {
	feud.BaseBarony
}

var BKalmius卡利米乌斯 feud.Barony = &卡利米乌斯KalmiusBarony{}

func init() {
    f := BKalmius卡利米乌斯.(*卡利米乌斯KalmiusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalmius",
		TitleName: "卡利米乌斯",
		TitleCode: "b_kalmius",
	}
}
