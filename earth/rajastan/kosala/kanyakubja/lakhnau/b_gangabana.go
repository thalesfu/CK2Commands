package lakhnau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恒伽婆那GangabanaBarony struct {
	feud.BaseBarony
}

var BGangabana恒伽婆那 feud.Barony = &恒伽婆那GangabanaBarony{}

func init() {
    f := BGangabana恒伽婆那.(*恒伽婆那GangabanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gangabana",
		TitleName: "恒伽婆那",
		TitleCode: "b_gangabana",
	}
}
