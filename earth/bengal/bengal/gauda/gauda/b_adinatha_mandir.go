package gauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿提那他神庙Adinatha_mandirBarony struct {
	feud.BaseBarony
}

var BAdinatha_mandir阿提那他神庙 feud.Barony = &阿提那他神庙Adinatha_mandirBarony{}

func init() {
    f := BAdinatha_mandir阿提那他神庙.(*阿提那他神庙Adinatha_mandirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adinatha_mandir",
		TitleName: "阿提那他神庙",
		TitleCode: "b_adinatha_mandir",
	}
}
