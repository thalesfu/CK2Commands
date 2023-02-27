package rama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉达查茨GradacacBarony struct {
	feud.BaseBarony
}

var BGradacac格拉达查茨 feud.Barony = &格拉达查茨GradacacBarony{}

func init() {
    f := BGradacac格拉达查茨.(*格拉达查茨GradacacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gradacac",
		TitleName: "格拉达查茨",
		TitleCode: "b_gradacac",
	}
}
