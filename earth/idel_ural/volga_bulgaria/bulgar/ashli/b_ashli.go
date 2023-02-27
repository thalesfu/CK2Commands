package ashli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿什勒AshliBarony struct {
	feud.BaseBarony
}

var BAshli阿什勒 feud.Barony = &阿什勒AshliBarony{}

func init() {
    f := BAshli阿什勒.(*阿什勒AshliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ashli",
		TitleName: "阿什勒",
		TitleCode: "b_ashli",
	}
}
