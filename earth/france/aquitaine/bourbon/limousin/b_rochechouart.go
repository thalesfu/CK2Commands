package limousin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗什舒瓦尔RochechouartBarony struct {
	feud.BaseBarony
}

var BRochechouart罗什舒瓦尔 feud.Barony = &罗什舒瓦尔RochechouartBarony{}

func init() {
    f := BRochechouart罗什舒瓦尔.(*罗什舒瓦尔RochechouartBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rochechouart",
		TitleName: "罗什舒瓦尔",
		TitleCode: "b_rochechouart",
	}
}
