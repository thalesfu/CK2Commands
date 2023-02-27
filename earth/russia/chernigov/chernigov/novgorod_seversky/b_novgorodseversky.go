package novgorod_seversky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺夫哥罗德_谢韦尔斯基NovgorodseverskyBarony struct {
	feud.BaseBarony
}

var BNovgorodseversky诺夫哥罗德_谢韦尔斯基 feud.Barony = &诺夫哥罗德_谢韦尔斯基NovgorodseverskyBarony{}

func init() {
    f := BNovgorodseversky诺夫哥罗德_谢韦尔斯基.(*诺夫哥罗德_谢韦尔斯基NovgorodseverskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novgorodseversky",
		TitleName: "诺夫哥罗德_谢韦尔斯基",
		TitleCode: "b_novgorodseversky",
	}
}
