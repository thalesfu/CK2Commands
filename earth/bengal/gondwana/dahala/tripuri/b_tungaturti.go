package tripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 东伽咄帝TungaturtiBarony struct {
	feud.BaseBarony
}

var BTungaturti东伽咄帝 feud.Barony = &东伽咄帝TungaturtiBarony{}

func init() {
	f := BTungaturti东伽咄帝.(*东伽咄帝TungaturtiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tungaturti",
		TitleName: "东伽咄帝",
		TitleCode: "b_tungaturti",
	}
}
