package madurai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿奈马莱AnaimalaiBarony struct {
	feud.BaseBarony
}

var BAnaimalai阿奈马莱 feud.Barony = &阿奈马莱AnaimalaiBarony{}

func init() {
    f := BAnaimalai阿奈马莱.(*阿奈马莱AnaimalaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anaimalai",
		TitleName: "阿奈马莱",
		TitleCode: "b_anaimalai",
	}
}
