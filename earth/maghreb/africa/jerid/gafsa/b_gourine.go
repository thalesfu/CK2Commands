package gafsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古林GourineBarony struct {
	feud.BaseBarony
}

var BGourine古林 feud.Barony = &古林GourineBarony{}

func init() {
    f := BGourine古林.(*古林GourineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gourine",
		TitleName: "古林",
		TitleCode: "b_gourine",
	}
}
