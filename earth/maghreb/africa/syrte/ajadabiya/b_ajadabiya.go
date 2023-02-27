package ajadabiya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾季达比耶AjadabiyaBarony struct {
	feud.BaseBarony
}

var BAjadabiya艾季达比耶 feud.Barony = &艾季达比耶AjadabiyaBarony{}

func init() {
    f := BAjadabiya艾季达比耶.(*艾季达比耶AjadabiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ajadabiya",
		TitleName: "艾季达比耶",
		TitleCode: "b_ajadabiya",
	}
}
