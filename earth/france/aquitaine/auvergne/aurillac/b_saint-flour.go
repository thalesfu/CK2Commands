package aurillac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣弗卢尔Saint_flourBarony struct {
	feud.BaseBarony
}

var BSaint_flour圣弗卢尔 feud.Barony = &圣弗卢尔Saint_flourBarony{}

func init() {
    f := BSaint_flour圣弗卢尔.(*圣弗卢尔Saint_flourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saint_flour",
		TitleName: "圣弗卢尔",
		TitleCode: "b_saint-flour",
	}
}
