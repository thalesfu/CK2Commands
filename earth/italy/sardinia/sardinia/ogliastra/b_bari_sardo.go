package ogliastra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴里萨尔多Bari_sardoBarony struct {
	feud.BaseBarony
}

var BBari_sardo巴里萨尔多 feud.Barony = &巴里萨尔多Bari_sardoBarony{}

func init() {
    f := BBari_sardo巴里萨尔多.(*巴里萨尔多Bari_sardoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bari_sardo",
		TitleName: "巴里萨尔多",
		TitleCode: "b_bari_sardo",
	}
}
