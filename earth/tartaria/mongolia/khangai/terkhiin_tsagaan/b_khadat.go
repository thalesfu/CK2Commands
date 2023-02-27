package terkhiin_tsagaan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈达特KhadatBarony struct {
	feud.BaseBarony
}

var BKhadat哈达特 feud.Barony = &哈达特KhadatBarony{}

func init() {
    f := BKhadat哈达特.(*哈达特KhadatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khadat",
		TitleName: "哈达特",
		TitleCode: "b_khadat",
	}
}
