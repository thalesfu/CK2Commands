package wadai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿贝歇AbecheBarony struct {
	feud.BaseBarony
}

var BAbeche阿贝歇 feud.Barony = &阿贝歇AbecheBarony{}

func init() {
	f := BAbeche阿贝歇.(*阿贝歇AbecheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abeche",
		TitleName: "阿贝歇",
		TitleCode: "b_abeche",
	}
}
