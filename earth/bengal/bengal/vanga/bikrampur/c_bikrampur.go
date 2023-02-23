package bikrampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BikrampurCounty interface {
	feud.County
	BBikrampur毗讫罗摩补罗() feud.Barony
	BHajiganj诃时健阇() feud.Barony
	BLoricol若利孔() feud.Barony
	BMeddappakkam摩陀波甘() feud.Barony
	BSutrapur修多罗补罗() feud.Barony
}

type 毗讫罗摩补罗BikrampurCounty struct {
	feud.BaseCounty
	毗讫罗摩补罗Bikrampur  feud.Barony
	诃时健阇Hajiganj     feud.Barony
	若利孔Loricol       feud.Barony
	摩陀波甘Meddappakkam feud.Barony
	修多罗补罗Sutrapur    feud.Barony
}

func (c *毗讫罗摩补罗BikrampurCounty) BBikrampur毗讫罗摩补罗() feud.Barony {
	return c.毗讫罗摩补罗Bikrampur
}

func (c *毗讫罗摩补罗BikrampurCounty) BHajiganj诃时健阇() feud.Barony {
	return c.诃时健阇Hajiganj
}

func (c *毗讫罗摩补罗BikrampurCounty) BLoricol若利孔() feud.Barony {
	return c.若利孔Loricol
}

func (c *毗讫罗摩补罗BikrampurCounty) BMeddappakkam摩陀波甘() feud.Barony {
	return c.摩陀波甘Meddappakkam
}

func (c *毗讫罗摩补罗BikrampurCounty) BSutrapur修多罗补罗() feud.Barony {
	return c.修多罗补罗Sutrapur
}

var CBikrampur毗讫罗摩补罗 BikrampurCounty = &毗讫罗摩补罗BikrampurCounty{}

func init() {
	f := CBikrampur毗讫罗摩补罗.(*毗讫罗摩补罗BikrampurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1319",
		Title:     "bikrampur",
		TitleName: "毗讫罗摩补罗",
		TitleCode: "c_bikrampur",
		Baronies:  map[string]feud.Barony{},
	}

	f.毗讫罗摩补罗Bikrampur = BBikrampur毗讫罗摩补罗
	f.毗讫罗摩补罗Bikrampur.SetParent(f)

	f.诃时健阇Hajiganj = BHajiganj诃时健阇
	f.诃时健阇Hajiganj.SetParent(f)

	f.若利孔Loricol = BLoricol若利孔
	f.若利孔Loricol.SetParent(f)

	f.摩陀波甘Meddappakkam = BMeddappakkam摩陀波甘
	f.摩陀波甘Meddappakkam.SetParent(f)

	f.修多罗补罗Sutrapur = BSutrapur修多罗补罗
	f.修多罗补罗Sutrapur.SetParent(f)

}
