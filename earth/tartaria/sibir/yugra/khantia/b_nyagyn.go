package khantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼亚甘NyagynBarony struct {
	feud.BaseBarony
}

var BNyagyn尼亚甘 feud.Barony = &尼亚甘NyagynBarony{}

func init() {
    f := BNyagyn尼亚甘.(*尼亚甘NyagynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyagyn",
		TitleName: "尼亚甘",
		TitleCode: "b_nyagyn",
	}
}
