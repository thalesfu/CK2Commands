package imotski

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉加伊BlagajBarony struct {
	feud.BaseBarony
}

var BBlagaj布拉加伊 feud.Barony = &布拉加伊BlagajBarony{}

func init() {
    f := BBlagaj布拉加伊.(*布拉加伊BlagajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "blagaj",
		TitleName: "布拉加伊",
		TitleCode: "b_blagaj",
	}
}
