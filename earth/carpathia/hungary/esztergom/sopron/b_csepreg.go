package sopron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柴普赖格CsepregBarony struct {
	feud.BaseBarony
}

var BCsepreg柴普赖格 feud.Barony = &柴普赖格CsepregBarony{}

func init() {
    f := BCsepreg柴普赖格.(*柴普赖格CsepregBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "csepreg",
		TitleName: "柴普赖格",
		TitleCode: "b_csepreg",
	}
}
