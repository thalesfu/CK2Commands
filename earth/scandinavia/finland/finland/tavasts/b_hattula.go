package tavasts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈图拉HattulaBarony struct {
	feud.BaseBarony
}

var BHattula哈图拉 feud.Barony = &哈图拉HattulaBarony{}

func init() {
    f := BHattula哈图拉.(*哈图拉HattulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hattula",
		TitleName: "哈图拉",
		TitleCode: "b_hattula",
	}
}
