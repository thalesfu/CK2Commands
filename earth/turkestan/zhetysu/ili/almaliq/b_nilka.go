package almaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼勒克NilkaBarony struct {
	feud.BaseBarony
}

var BNilka尼勒克 feud.Barony = &尼勒克NilkaBarony{}

func init() {
    f := BNilka尼勒克.(*尼勒克NilkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nilka",
		TitleName: "尼勒克",
		TitleCode: "b_nilka",
	}
}
