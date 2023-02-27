package plauen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克诺贝尔斯多夫KnobelsdorfBarony struct {
	feud.BaseBarony
}

var BKnobelsdorf克诺贝尔斯多夫 feud.Barony = &克诺贝尔斯多夫KnobelsdorfBarony{}

func init() {
    f := BKnobelsdorf克诺贝尔斯多夫.(*克诺贝尔斯多夫KnobelsdorfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "knobelsdorf",
		TitleName: "克诺贝尔斯多夫",
		TitleCode: "b_knobelsdorf",
	}
}
