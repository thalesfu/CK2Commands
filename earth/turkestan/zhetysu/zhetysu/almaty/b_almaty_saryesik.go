package almaty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉木图Almaty_saryesikBarony struct {
	feud.BaseBarony
}

var BAlmaty_saryesik阿拉木图 feud.Barony = &阿拉木图Almaty_saryesikBarony{}

func init() {
    f := BAlmaty_saryesik阿拉木图.(*阿拉木图Almaty_saryesikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almaty_saryesik",
		TitleName: "阿拉木图",
		TitleCode: "b_almaty_saryesik",
	}
}
