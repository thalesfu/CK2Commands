package wight

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡里斯布鲁克CarisbrookeBarony struct {
	feud.BaseBarony
}

var BCarisbrooke卡里斯布鲁克 feud.Barony = &卡里斯布鲁克CarisbrookeBarony{}

func init() {
    f := BCarisbrooke卡里斯布鲁克.(*卡里斯布鲁克CarisbrookeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carisbrooke",
		TitleName: "卡里斯布鲁克",
		TitleCode: "b_carisbrooke",
	}
}
