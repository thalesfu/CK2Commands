package burgos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜罗河畔阿兰达ArandadedueroBarony struct {
	feud.BaseBarony
}

var BArandadeduero杜罗河畔阿兰达 feud.Barony = &杜罗河畔阿兰达ArandadedueroBarony{}

func init() {
    f := BArandadeduero杜罗河畔阿兰达.(*杜罗河畔阿兰达ArandadedueroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arandadeduero",
		TitleName: "杜罗河畔阿兰达",
		TitleCode: "b_arandadeduero",
	}
}
