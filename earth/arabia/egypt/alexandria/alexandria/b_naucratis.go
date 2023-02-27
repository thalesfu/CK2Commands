package alexandria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺克拉提斯NaucratisBarony struct {
	feud.BaseBarony
}

var BNaucratis诺克拉提斯 feud.Barony = &诺克拉提斯NaucratisBarony{}

func init() {
    f := BNaucratis诺克拉提斯.(*诺克拉提斯NaucratisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naucratis",
		TitleName: "诺克拉提斯",
		TitleCode: "b_naucratis",
	}
}
