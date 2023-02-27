package niederbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡姆ChamBarony struct {
	feud.BaseBarony
}

var BCham卡姆 feud.Barony = &卡姆ChamBarony{}

func init() {
    f := BCham卡姆.(*卡姆ChamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cham",
		TitleName: "卡姆",
		TitleCode: "b_cham",
	}
}
