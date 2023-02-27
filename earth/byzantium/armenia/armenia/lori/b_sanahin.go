package lori

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨那欣SanahinBarony struct {
	feud.BaseBarony
}

var BSanahin萨那欣 feud.Barony = &萨那欣SanahinBarony{}

func init() {
    f := BSanahin萨那欣.(*萨那欣SanahinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanahin",
		TitleName: "萨那欣",
		TitleCode: "b_sanahin",
	}
}
