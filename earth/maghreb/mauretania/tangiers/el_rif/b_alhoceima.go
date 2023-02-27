package el_rif

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡塞马AlhoceimaBarony struct {
	feud.BaseBarony
}

var BAlhoceima胡塞马 feud.Barony = &胡塞马AlhoceimaBarony{}

func init() {
    f := BAlhoceima胡塞马.(*胡塞马AlhoceimaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alhoceima",
		TitleName: "胡塞马",
		TitleCode: "b_alhoceima",
	}
}
