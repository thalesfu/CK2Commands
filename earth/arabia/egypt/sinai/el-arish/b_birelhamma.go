package el-arish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈迈井BirelhammaBarony struct {
	feud.BaseBarony
}

var BBirelhamma哈迈井 feud.Barony = &哈迈井BirelhammaBarony{}

func init() {
    f := BBirelhamma哈迈井.(*哈迈井BirelhammaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birelhamma",
		TitleName: "哈迈井",
		TitleCode: "b_birelhamma",
	}
}
