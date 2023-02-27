package lorraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣阿沃尔德SaintavoidBarony struct {
	feud.BaseBarony
}

var BSaintavoid圣阿沃尔德 feud.Barony = &圣阿沃尔德SaintavoidBarony{}

func init() {
    f := BSaintavoid圣阿沃尔德.(*圣阿沃尔德SaintavoidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintavoid",
		TitleName: "圣阿沃尔德",
		TitleCode: "b_saintavoid",
	}
}
