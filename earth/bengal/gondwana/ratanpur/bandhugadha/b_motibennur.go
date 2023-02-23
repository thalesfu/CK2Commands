package bandhugadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 牟提旁奴MotibennurBarony struct {
	feud.BaseBarony
}

var BMotibennur牟提旁奴 feud.Barony = &牟提旁奴MotibennurBarony{}

func init() {
	f := BMotibennur牟提旁奴.(*牟提旁奴MotibennurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "motibennur",
		TitleName: "牟提旁奴",
		TitleCode: "b_motibennur",
	}
}
