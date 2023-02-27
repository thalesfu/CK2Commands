package opava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩拉维采河畔赫拉德茨Hradec_nad_moraviciBarony struct {
	feud.BaseBarony
}

var BHradec_nad_moravici摩拉维采河畔赫拉德茨 feud.Barony = &摩拉维采河畔赫拉德茨Hradec_nad_moraviciBarony{}

func init() {
    f := BHradec_nad_moravici摩拉维采河畔赫拉德茨.(*摩拉维采河畔赫拉德茨Hradec_nad_moraviciBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hradec_nad_moravici",
		TitleName: "摩拉维采河畔赫拉德茨",
		TitleCode: "b_hradec_nad_moravici",
	}
}
