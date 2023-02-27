package rohana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迭伽婆比DighavapiBarony struct {
	feud.BaseBarony
}

var BDighavapi迭伽婆比 feud.Barony = &迭伽婆比DighavapiBarony{}

func init() {
    f := BDighavapi迭伽婆比.(*迭伽婆比DighavapiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dighavapi",
		TitleName: "迭伽婆比",
		TitleCode: "b_dighavapi",
	}
}
