package kostroma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏季斯拉夫尔SudislavlBarony struct {
	feud.BaseBarony
}

var BSudislavl苏季斯拉夫尔 feud.Barony = &苏季斯拉夫尔SudislavlBarony{}

func init() {
    f := BSudislavl苏季斯拉夫尔.(*苏季斯拉夫尔SudislavlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sudislavl",
		TitleName: "苏季斯拉夫尔",
		TitleCode: "b_sudislavl",
	}
}
