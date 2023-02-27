package chanderi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦尔贡陀尔Garh_kundarBarony struct {
	feud.BaseBarony
}

var BGarh_kundar迦尔贡陀尔 feud.Barony = &迦尔贡陀尔Garh_kundarBarony{}

func init() {
    f := BGarh_kundar迦尔贡陀尔.(*迦尔贡陀尔Garh_kundarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "garh_kundar",
		TitleName: "迦尔贡陀尔",
		TitleCode: "b_garh_kundar",
	}
}
