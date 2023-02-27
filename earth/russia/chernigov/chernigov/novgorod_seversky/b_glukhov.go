package novgorod_seversky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格卢霍夫GlukhovBarony struct {
	feud.BaseBarony
}

var BGlukhov格卢霍夫 feud.Barony = &格卢霍夫GlukhovBarony{}

func init() {
    f := BGlukhov格卢霍夫.(*格卢霍夫GlukhovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glukhov",
		TitleName: "格卢霍夫",
		TitleCode: "b_glukhov",
	}
}
