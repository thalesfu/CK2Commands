package pereyaslavl_zalessky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉巴诺沃KarabanovoBarony struct {
	feud.BaseBarony
}

var BKarabanovo卡拉巴诺沃 feud.Barony = &卡拉巴诺沃KarabanovoBarony{}

func init() {
    f := BKarabanovo卡拉巴诺沃.(*卡拉巴诺沃KarabanovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karabanovo",
		TitleName: "卡拉巴诺沃",
		TitleCode: "b_karabanovo",
	}
}
