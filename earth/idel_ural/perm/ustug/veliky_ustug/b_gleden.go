package veliky_ustug

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格列坚GledenBarony struct {
	feud.BaseBarony
}

var BGleden格列坚 feud.Barony = &格列坚GledenBarony{}

func init() {
    f := BGleden格列坚.(*格列坚GledenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gleden",
		TitleName: "格列坚",
		TitleCode: "b_gleden",
	}
}
