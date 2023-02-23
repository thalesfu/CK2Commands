package markam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 中林ZhonglinBarony struct {
	feud.BaseBarony
}

var BZhonglin中林 feud.Barony = &中林ZhonglinBarony{}

func init() {
	f := BZhonglin中林.(*中林ZhonglinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhonglin",
		TitleName: "中林",
		TitleCode: "b_zhonglin",
	}
}
