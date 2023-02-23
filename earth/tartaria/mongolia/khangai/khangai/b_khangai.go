package khangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杭海KhangaiBarony struct {
	feud.BaseBarony
}

var BKhangai杭海 feud.Barony = &杭海KhangaiBarony{}

func init() {
	f := BKhangai杭海.(*杭海KhangaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khangai",
		TitleName: "杭海",
		TitleCode: "b_khangai",
	}
}
