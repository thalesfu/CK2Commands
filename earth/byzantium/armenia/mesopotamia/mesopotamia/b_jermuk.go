package mesopotamia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉尔穆克JermukBarony struct {
	feud.BaseBarony
}

var BJermuk吉尔穆克 feud.Barony = &吉尔穆克JermukBarony{}

func init() {
	f := BJermuk吉尔穆克.(*吉尔穆克JermukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jermuk",
		TitleName: "吉尔穆克",
		TitleCode: "b_jermuk",
	}
}
