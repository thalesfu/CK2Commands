package alexandretta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉罗什德里索勒LarochederissoleBarony struct {
	feud.BaseBarony
}

var BLarochederissole拉罗什德里索勒 feud.Barony = &拉罗什德里索勒LarochederissoleBarony{}

func init() {
	f := BLarochederissole拉罗什德里索勒.(*拉罗什德里索勒LarochederissoleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "larochederissole",
		TitleName: "拉罗什德里索勒",
		TitleCode: "b_larochederissole",
	}
}
