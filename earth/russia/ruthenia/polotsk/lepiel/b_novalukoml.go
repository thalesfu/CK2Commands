package lepiel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺瓦卢科姆利NovalukomlBarony struct {
	feud.BaseBarony
}

var BNovalukoml诺瓦卢科姆利 feud.Barony = &诺瓦卢科姆利NovalukomlBarony{}

func init() {
	f := BNovalukoml诺瓦卢科姆利.(*诺瓦卢科姆利NovalukomlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novalukoml",
		TitleName: "诺瓦卢科姆利",
		TitleCode: "b_novalukoml",
	}
}
