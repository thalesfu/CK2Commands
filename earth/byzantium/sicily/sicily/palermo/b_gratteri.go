package palermo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉泰里GratteriBarony struct {
	feud.BaseBarony
}

var BGratteri格拉泰里 feud.Barony = &格拉泰里GratteriBarony{}

func init() {
    f := BGratteri格拉泰里.(*格拉泰里GratteriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gratteri",
		TitleName: "格拉泰里",
		TitleCode: "b_gratteri",
	}
}
