package niederbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基希罗特KirchrothBarony struct {
	feud.BaseBarony
}

var BKirchroth基希罗特 feud.Barony = &基希罗特KirchrothBarony{}

func init() {
    f := BKirchroth基希罗特.(*基希罗特KirchrothBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirchroth",
		TitleName: "基希罗特",
		TitleCode: "b_kirchroth",
	}
}
