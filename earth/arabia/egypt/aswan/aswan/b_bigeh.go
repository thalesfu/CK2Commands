package aswan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比格BigehBarony struct {
	feud.BaseBarony
}

var BBigeh比格 feud.Barony = &比格BigehBarony{}

func init() {
	f := BBigeh比格.(*比格BigehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bigeh",
		TitleName: "比格",
		TitleCode: "b_bigeh",
	}
}
