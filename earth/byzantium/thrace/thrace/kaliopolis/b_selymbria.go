package kaliopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟利姆布里亚SelymbriaBarony struct {
	feud.BaseBarony
}

var BSelymbria瑟利姆布里亚 feud.Barony = &瑟利姆布里亚SelymbriaBarony{}

func init() {
	f := BSelymbria瑟利姆布里亚.(*瑟利姆布里亚SelymbriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "selymbria",
		TitleName: "瑟利姆布里亚",
		TitleCode: "b_selymbria",
	}
}
