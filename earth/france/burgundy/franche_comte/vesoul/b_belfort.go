package vesoul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔福BelfortBarony struct {
	feud.BaseBarony
}

var BBelfort贝尔福 feud.Barony = &贝尔福BelfortBarony{}

func init() {
    f := BBelfort贝尔福.(*贝尔福BelfortBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belfort",
		TitleName: "贝尔福",
		TitleCode: "b_belfort",
	}
}
