package hajr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 荷台达AlhudaydahBarony struct {
	feud.BaseBarony
}

var BAlhudaydah荷台达 feud.Barony = &荷台达AlhudaydahBarony{}

func init() {
    f := BAlhudaydah荷台达.(*荷台达AlhudaydahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alhudaydah",
		TitleName: "荷台达",
		TitleCode: "b_alhudaydah",
	}
}
