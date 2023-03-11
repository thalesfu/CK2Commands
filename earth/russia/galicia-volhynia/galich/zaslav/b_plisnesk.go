package zaslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普列斯涅斯克PlisneskBarony struct {
	feud.BaseBarony
}

var BPlisnesk普列斯涅斯克 feud.Barony = &普列斯涅斯克PlisneskBarony{}

func init() {
    f := BPlisnesk普列斯涅斯克.(*普列斯涅斯克PlisneskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plisnesk",
		TitleName: "普列斯涅斯克",
		TitleCode: "b_plisnesk",
	}
}
