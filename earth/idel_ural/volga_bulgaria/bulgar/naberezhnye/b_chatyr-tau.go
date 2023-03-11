package naberezhnye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰特尔_陶Chatyr_tauBarony struct {
	feud.BaseBarony
}

var BChatyr_tau恰特尔_陶 feud.Barony = &恰特尔_陶Chatyr_tauBarony{}

func init() {
    f := BChatyr_tau恰特尔_陶.(*恰特尔_陶Chatyr_tauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chatyr_tau",
		TitleName: "恰特尔_陶",
		TitleCode: "b_chatyr-tau",
	}
}
