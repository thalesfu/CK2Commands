package poznanskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰恩库夫CzarnkowBarony struct {
	feud.BaseBarony
}

var BCzarnkow恰恩库夫 feud.Barony = &恰恩库夫CzarnkowBarony{}

func init() {
    f := BCzarnkow恰恩库夫.(*恰恩库夫CzarnkowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "czarnkow",
		TitleName: "恰恩库夫",
		TitleCode: "b_czarnkow",
	}
}
