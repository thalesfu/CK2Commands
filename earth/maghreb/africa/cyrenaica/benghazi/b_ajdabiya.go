package benghazi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾季达比耶AjdabiyaBarony struct {
	feud.BaseBarony
}

var BAjdabiya艾季达比耶 feud.Barony = &艾季达比耶AjdabiyaBarony{}

func init() {
    f := BAjdabiya艾季达比耶.(*艾季达比耶AjdabiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ajdabiya",
		TitleName: "艾季达比耶",
		TitleCode: "b_ajdabiya",
	}
}
