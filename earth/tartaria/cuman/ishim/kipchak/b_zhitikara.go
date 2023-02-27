package kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰特加拉ZhitikaraBarony struct {
	feud.BaseBarony
}

var BZhitikara杰特加拉 feud.Barony = &杰特加拉ZhitikaraBarony{}

func init() {
    f := BZhitikara杰特加拉.(*杰特加拉ZhitikaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhitikara",
		TitleName: "杰特加拉",
		TitleCode: "b_zhitikara",
	}
}
