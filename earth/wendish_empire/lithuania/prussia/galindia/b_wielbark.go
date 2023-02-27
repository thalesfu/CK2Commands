package galindia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尔巴克WielbarkBarony struct {
	feud.BaseBarony
}

var BWielbark维尔巴克 feud.Barony = &维尔巴克WielbarkBarony{}

func init() {
    f := BWielbark维尔巴克.(*维尔巴克WielbarkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wielbark",
		TitleName: "维尔巴克",
		TitleCode: "b_wielbark",
	}
}
