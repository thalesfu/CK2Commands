package somerset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汤顿TauntonBarony struct {
	feud.BaseBarony
}

var BTaunton汤顿 feud.Barony = &汤顿TauntonBarony{}

func init() {
	f := BTaunton汤顿.(*汤顿TauntonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taunton",
		TitleName: "汤顿",
		TitleCode: "b_taunton",
	}
}
