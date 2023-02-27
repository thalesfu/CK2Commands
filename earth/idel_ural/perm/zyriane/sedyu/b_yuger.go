package sedyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤格尔YugerBarony struct {
	feud.BaseBarony
}

var BYuger尤格尔 feud.Barony = &尤格尔YugerBarony{}

func init() {
    f := BYuger尤格尔.(*尤格尔YugerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yuger",
		TitleName: "尤格尔",
		TitleCode: "b_yuger",
	}
}
