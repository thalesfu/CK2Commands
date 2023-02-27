package imotski

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格鲁德GrudeBarony struct {
	feud.BaseBarony
}

var BGrude格鲁德 feud.Barony = &格鲁德GrudeBarony{}

func init() {
    f := BGrude格鲁德.(*格鲁德GrudeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grude",
		TitleName: "格鲁德",
		TitleCode: "b_grude",
	}
}
