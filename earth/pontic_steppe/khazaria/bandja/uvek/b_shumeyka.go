package uvek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒梅卡ShumeykaBarony struct {
	feud.BaseBarony
}

var BShumeyka舒梅卡 feud.Barony = &舒梅卡ShumeykaBarony{}

func init() {
    f := BShumeyka舒梅卡.(*舒梅卡ShumeykaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shumeyka",
		TitleName: "舒梅卡",
		TitleCode: "b_shumeyka",
	}
}
