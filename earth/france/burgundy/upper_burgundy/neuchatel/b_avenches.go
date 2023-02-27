package neuchatel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿旺什AvenchesBarony struct {
	feud.BaseBarony
}

var BAvenches阿旺什 feud.Barony = &阿旺什AvenchesBarony{}

func init() {
    f := BAvenches阿旺什.(*阿旺什AvenchesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avenches",
		TitleName: "阿旺什",
		TitleCode: "b_avenches",
	}
}
