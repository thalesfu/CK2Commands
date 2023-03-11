package belz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔兹BelzBarony struct {
	feud.BaseBarony
}

var BBelz贝尔兹 feud.Barony = &贝尔兹BelzBarony{}

func init() {
    f := BBelz贝尔兹.(*贝尔兹BelzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belz",
		TitleName: "贝尔兹",
		TitleCode: "b_belz",
	}
}
