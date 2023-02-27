package lancaster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克利瑟罗ClitheroeBarony struct {
	feud.BaseBarony
}

var BClitheroe克利瑟罗 feud.Barony = &克利瑟罗ClitheroeBarony{}

func init() {
    f := BClitheroe克利瑟罗.(*克利瑟罗ClitheroeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clitheroe",
		TitleName: "克利瑟罗",
		TitleCode: "b_clitheroe",
	}
}
