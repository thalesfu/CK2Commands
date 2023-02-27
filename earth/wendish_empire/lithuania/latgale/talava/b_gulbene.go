package talava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔贝内GulbeneBarony struct {
	feud.BaseBarony
}

var BGulbene古尔贝内 feud.Barony = &古尔贝内GulbeneBarony{}

func init() {
    f := BGulbene古尔贝内.(*古尔贝内GulbeneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gulbene",
		TitleName: "古尔贝内",
		TitleCode: "b_gulbene",
	}
}
