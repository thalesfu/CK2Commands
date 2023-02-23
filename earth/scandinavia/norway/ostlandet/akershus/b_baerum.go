package akershus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜鲁姆BaerumBarony struct {
	feud.BaseBarony
}

var BBaerum拜鲁姆 feud.Barony = &拜鲁姆BaerumBarony{}

func init() {
	f := BBaerum拜鲁姆.(*拜鲁姆BaerumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baerum",
		TitleName: "拜鲁姆",
		TitleCode: "b_baerum",
	}
}
