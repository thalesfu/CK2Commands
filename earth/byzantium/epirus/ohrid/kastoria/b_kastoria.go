package kastoria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯托里亚KastoriaBarony struct {
	feud.BaseBarony
}

var BKastoria卡斯托里亚 feud.Barony = &卡斯托里亚KastoriaBarony{}

func init() {
	f := BKastoria卡斯托里亚.(*卡斯托里亚KastoriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kastoria",
		TitleName: "卡斯托里亚",
		TitleCode: "b_kastoria",
	}
}
