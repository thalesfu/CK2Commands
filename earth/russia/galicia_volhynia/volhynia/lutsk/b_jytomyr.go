package lutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日托米尔JytomyrBarony struct {
	feud.BaseBarony
}

var BJytomyr日托米尔 feud.Barony = &日托米尔JytomyrBarony{}

func init() {
    f := BJytomyr日托米尔.(*日托米尔JytomyrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jytomyr",
		TitleName: "日托米尔",
		TitleCode: "b_jytomyr",
	}
}
