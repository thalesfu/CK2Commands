package aquileia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉多GradoBarony struct {
	feud.BaseBarony
}

var BGrado格拉多 feud.Barony = &格拉多GradoBarony{}

func init() {
    f := BGrado格拉多.(*格拉多GradoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grado",
		TitleName: "格拉多",
		TitleCode: "b_grado",
	}
}
