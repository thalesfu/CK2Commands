package tortosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦拉尼亚ValaniaBarony struct {
	feud.BaseBarony
}

var BValania瓦拉尼亚 feud.Barony = &瓦拉尼亚ValaniaBarony{}

func init() {
    f := BValania瓦拉尼亚.(*瓦拉尼亚ValaniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valania",
		TitleName: "瓦拉尼亚",
		TitleCode: "b_valania",
	}
}
