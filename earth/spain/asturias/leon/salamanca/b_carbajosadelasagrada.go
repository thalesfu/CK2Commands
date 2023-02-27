package salamanca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔瓦霍萨德拉萨格拉达CarbajosadelasagradaBarony struct {
	feud.BaseBarony
}

var BCarbajosadelasagrada卡尔瓦霍萨德拉萨格拉达 feud.Barony = &卡尔瓦霍萨德拉萨格拉达CarbajosadelasagradaBarony{}

func init() {
    f := BCarbajosadelasagrada卡尔瓦霍萨德拉萨格拉达.(*卡尔瓦霍萨德拉萨格拉达CarbajosadelasagradaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carbajosadelasagrada",
		TitleName: "卡尔瓦霍萨德拉萨格拉达",
		TitleCode: "b_carbajosadelasagrada",
	}
}
