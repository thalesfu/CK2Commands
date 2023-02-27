package north_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科里亚日马KoryazhmaBarony struct {
	feud.BaseBarony
}

var BKoryazhma科里亚日马 feud.Barony = &科里亚日马KoryazhmaBarony{}

func init() {
    f := BKoryazhma科里亚日马.(*科里亚日马KoryazhmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koryazhma",
		TitleName: "科里亚日马",
		TitleCode: "b_koryazhma",
	}
}
