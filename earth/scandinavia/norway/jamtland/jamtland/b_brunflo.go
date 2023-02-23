package jamtland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布伦弗卢BrunfloBarony struct {
	feud.BaseBarony
}

var BBrunflo布伦弗卢 feud.Barony = &布伦弗卢BrunfloBarony{}

func init() {
	f := BBrunflo布伦弗卢.(*布伦弗卢BrunfloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brunflo",
		TitleName: "布伦弗卢",
		TitleCode: "b_brunflo",
	}
}
