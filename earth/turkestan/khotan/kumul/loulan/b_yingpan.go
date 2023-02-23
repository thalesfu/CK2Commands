package loulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 营盘YingpanBarony struct {
	feud.BaseBarony
}

var BYingpan营盘 feud.Barony = &营盘YingpanBarony{}

func init() {
	f := BYingpan营盘.(*营盘YingpanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yingpan",
		TitleName: "营盘",
		TitleCode: "b_yingpan",
	}
}
