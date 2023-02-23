package vladimir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗日杰斯特文斯科耶RochdestvenskoyeBarony struct {
	feud.BaseBarony
}

var BRochdestvenskoye罗日杰斯特文斯科耶 feud.Barony = &罗日杰斯特文斯科耶RochdestvenskoyeBarony{}

func init() {
	f := BRochdestvenskoye罗日杰斯特文斯科耶.(*罗日杰斯特文斯科耶RochdestvenskoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rochdestvenskoye",
		TitleName: "罗日杰斯特文斯科耶",
		TitleCode: "b_rochdestvenskoye",
	}
}
