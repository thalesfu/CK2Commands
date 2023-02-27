package faiyum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尼苏埃夫Beni_suefBarony struct {
	feud.BaseBarony
}

var BBeni_suef贝尼苏埃夫 feud.Barony = &贝尼苏埃夫Beni_suefBarony{}

func init() {
    f := BBeni_suef贝尼苏埃夫.(*贝尼苏埃夫Beni_suefBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beni_suef",
		TitleName: "贝尼苏埃夫",
		TitleCode: "b_beni_suef",
	}
}
