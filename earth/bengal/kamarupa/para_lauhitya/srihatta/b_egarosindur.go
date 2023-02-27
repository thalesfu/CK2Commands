package srihatta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夷迦卢斯那突EgarosindurBarony struct {
	feud.BaseBarony
}

var BEgarosindur夷迦卢斯那突 feud.Barony = &夷迦卢斯那突EgarosindurBarony{}

func init() {
    f := BEgarosindur夷迦卢斯那突.(*夷迦卢斯那突EgarosindurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egarosindur",
		TitleName: "夷迦卢斯那突",
		TitleCode: "b_egarosindur",
	}
}
