package pannagallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳加尔库尔诺奥尔NagarkurnoolBarony struct {
	feud.BaseBarony
}

var BNagarkurnool纳加尔库尔诺奥尔 feud.Barony = &纳加尔库尔诺奥尔NagarkurnoolBarony{}

func init() {
    f := BNagarkurnool纳加尔库尔诺奥尔.(*纳加尔库尔诺奥尔NagarkurnoolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagarkurnool",
		TitleName: "纳加尔库尔诺奥尔",
		TitleCode: "b_nagarkurnool",
	}
}
