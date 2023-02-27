package mahdia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆萨肯MsakenBarony struct {
	feud.BaseBarony
}

var BMsaken姆萨肯 feud.Barony = &姆萨肯MsakenBarony{}

func init() {
    f := BMsaken姆萨肯.(*姆萨肯MsakenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "msaken",
		TitleName: "姆萨肯",
		TitleCode: "b_msaken",
	}
}
