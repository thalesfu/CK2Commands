package bezichy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尼亚日哈KnyazhikhaBarony struct {
	feud.BaseBarony
}

var BKnyazhikha克尼亚日哈 feud.Barony = &克尼亚日哈KnyazhikhaBarony{}

func init() {
    f := BKnyazhikha克尼亚日哈.(*克尼亚日哈KnyazhikhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "knyazhikha",
		TitleName: "克尼亚日哈",
		TitleCode: "b_knyazhikha",
	}
}
