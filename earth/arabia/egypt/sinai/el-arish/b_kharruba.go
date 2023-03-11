package el-arish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈鲁巴KharrubaBarony struct {
	feud.BaseBarony
}

var BKharruba哈鲁巴 feud.Barony = &哈鲁巴KharrubaBarony{}

func init() {
    f := BKharruba哈鲁巴.(*哈鲁巴KharrubaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kharruba",
		TitleName: "哈鲁巴",
		TitleCode: "b_kharruba",
	}
}
