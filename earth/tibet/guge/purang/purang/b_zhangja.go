package purang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 生结ZhangjaBarony struct {
	feud.BaseBarony
}

var BZhangja生结 feud.Barony = &生结ZhangjaBarony{}

func init() {
	f := BZhangja生结.(*生结ZhangjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhangja",
		TitleName: "生结",
		TitleCode: "b_zhangja",
	}
}
