package veglia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茨里克韦尼察CrikvenicaBarony struct {
	feud.BaseBarony
}

var BCrikvenica茨里克韦尼察 feud.Barony = &茨里克韦尼察CrikvenicaBarony{}

func init() {
    f := BCrikvenica茨里克韦尼察.(*茨里克韦尼察CrikvenicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crikvenica",
		TitleName: "茨里克韦尼察",
		TitleCode: "b_crikvenica",
	}
}
