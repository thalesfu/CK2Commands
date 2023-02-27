package brandenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔齐希BelzigBarony struct {
	feud.BaseBarony
}

var BBelzig贝尔齐希 feud.Barony = &贝尔齐希BelzigBarony{}

func init() {
    f := BBelzig贝尔齐希.(*贝尔齐希BelzigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belzig",
		TitleName: "贝尔齐希",
		TitleCode: "b_belzig",
	}
}
