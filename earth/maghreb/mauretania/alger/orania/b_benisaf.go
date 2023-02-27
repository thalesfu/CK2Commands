package orania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尼萨夫BenisafBarony struct {
	feud.BaseBarony
}

var BBenisaf贝尼萨夫 feud.Barony = &贝尼萨夫BenisafBarony{}

func init() {
    f := BBenisaf贝尼萨夫.(*贝尼萨夫BenisafBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "benisaf",
		TitleName: "贝尼萨夫",
		TitleCode: "b_benisaf",
	}
}
