package godwad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗罗摩罗BhillamalaBarony struct {
	feud.BaseBarony
}

var BBhillamala毗罗摩罗 feud.Barony = &毗罗摩罗BhillamalaBarony{}

func init() {
    f := BBhillamala毗罗摩罗.(*毗罗摩罗BhillamalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhillamala",
		TitleName: "毗罗摩罗",
		TitleCode: "b_bhillamala",
	}
}
