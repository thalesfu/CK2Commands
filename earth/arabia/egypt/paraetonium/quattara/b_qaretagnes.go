package quattara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀雷特阿格尼斯QaretagnesBarony struct {
	feud.BaseBarony
}

var BQaretagnes喀雷特阿格尼斯 feud.Barony = &喀雷特阿格尼斯QaretagnesBarony{}

func init() {
	f := BQaretagnes喀雷特阿格尼斯.(*喀雷特阿格尼斯QaretagnesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qaretagnes",
		TitleName: "喀雷特阿格尼斯",
		TitleCode: "b_qaretagnes",
	}
}
