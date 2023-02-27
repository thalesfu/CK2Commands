package tell_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳斯那NacinaBarony struct {
	feud.BaseBarony
}

var BNacina纳斯那 feud.Barony = &纳斯那NacinaBarony{}

func init() {
    f := BNacina纳斯那.(*纳斯那NacinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nacina",
		TitleName: "纳斯那",
		TitleCode: "b_nacina",
	}
}
