package nyingchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布久色吉拉康Buchasergyi_lakangBarony struct {
	feud.BaseBarony
}

var BBuchasergyi_lakang布久色吉拉康 feud.Barony = &布久色吉拉康Buchasergyi_lakangBarony{}

func init() {
    f := BBuchasergyi_lakang布久色吉拉康.(*布久色吉拉康Buchasergyi_lakangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buchasergyi_lakang",
		TitleName: "布久色吉拉康",
		TitleCode: "b_buchasergyi_lakang",
	}
}
