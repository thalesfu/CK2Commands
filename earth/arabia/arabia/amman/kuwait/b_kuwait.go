package kuwait

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科威特KuwaitBarony struct {
	feud.BaseBarony
}

var BKuwait科威特 feud.Barony = &科威特KuwaitBarony{}

func init() {
    f := BKuwait科威特.(*科威特KuwaitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuwait",
		TitleName: "科威特",
		TitleCode: "b_kuwait",
	}
}
