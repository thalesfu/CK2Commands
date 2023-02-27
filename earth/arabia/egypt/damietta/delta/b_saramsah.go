package delta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞慕斯SaramsahBarony struct {
	feud.BaseBarony
}

var BSaramsah塞慕斯 feud.Barony = &塞慕斯SaramsahBarony{}

func init() {
    f := BSaramsah塞慕斯.(*塞慕斯SaramsahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saramsah",
		TitleName: "塞慕斯",
		TitleCode: "b_saramsah",
	}
}
