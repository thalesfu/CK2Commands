package jerusalem

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝特努巴BeitnubaBarony struct {
	feud.BaseBarony
}

var BBeitnuba贝特努巴 feud.Barony = &贝特努巴BeitnubaBarony{}

func init() {
    f := BBeitnuba贝特努巴.(*贝特努巴BeitnubaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beitnuba",
		TitleName: "贝特努巴",
		TitleCode: "b_beitnuba",
	}
}
