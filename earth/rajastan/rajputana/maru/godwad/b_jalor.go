package godwad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦卢尔JalorBarony struct {
	feud.BaseBarony
}

var BJalor迦卢尔 feud.Barony = &迦卢尔JalorBarony{}

func init() {
	f := BJalor迦卢尔.(*迦卢尔JalorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jalor",
		TitleName: "迦卢尔",
		TitleCode: "b_jalor",
	}
}
