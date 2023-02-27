package kaneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克洛汀AkrotinBarony struct {
	feud.BaseBarony
}

var BAkrotin阿克洛汀 feud.Barony = &阿克洛汀AkrotinBarony{}

func init() {
    f := BAkrotin阿克洛汀.(*阿克洛汀AkrotinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akrotin",
		TitleName: "阿克洛汀",
		TitleCode: "b_akrotin",
	}
}
