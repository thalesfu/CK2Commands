package hormuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴姆BamBarony struct {
	feud.BaseBarony
}

var BBam巴姆 feud.Barony = &巴姆BamBarony{}

func init() {
    f := BBam巴姆.(*巴姆BamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bam",
		TitleName: "巴姆",
		TitleCode: "b_bam",
	}
}
