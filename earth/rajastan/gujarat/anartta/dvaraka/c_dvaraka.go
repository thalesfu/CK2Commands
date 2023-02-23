package dvaraka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DvarakaCounty interface {
	feud.County
	BBhanvad般跋() feud.Barony
	BDvaraka堕罗迦() feud.Barony
	BDwarakadheesh堕罗迦地舍() feud.Barony
	BGomati戈摩蒂() feud.Barony
	BHuljanti呼旬提() feud.Barony
	BJakar阇羯罗() feud.Barony
	BLalpur那尔补罗() feud.Barony
}

type 堕罗迦DvarakaCounty struct {
	feud.BaseCounty
	般跋Bhanvad          feud.Barony
	堕罗迦Dvaraka         feud.Barony
	堕罗迦地舍Dwarakadheesh feud.Barony
	戈摩蒂Gomati          feud.Barony
	呼旬提Huljanti        feud.Barony
	阇羯罗Jakar           feud.Barony
	那尔补罗Lalpur         feud.Barony
}

func (c *堕罗迦DvarakaCounty) BBhanvad般跋() feud.Barony {
	return c.般跋Bhanvad
}

func (c *堕罗迦DvarakaCounty) BDvaraka堕罗迦() feud.Barony {
	return c.堕罗迦Dvaraka
}

func (c *堕罗迦DvarakaCounty) BDwarakadheesh堕罗迦地舍() feud.Barony {
	return c.堕罗迦地舍Dwarakadheesh
}

func (c *堕罗迦DvarakaCounty) BGomati戈摩蒂() feud.Barony {
	return c.戈摩蒂Gomati
}

func (c *堕罗迦DvarakaCounty) BHuljanti呼旬提() feud.Barony {
	return c.呼旬提Huljanti
}

func (c *堕罗迦DvarakaCounty) BJakar阇羯罗() feud.Barony {
	return c.阇羯罗Jakar
}

func (c *堕罗迦DvarakaCounty) BLalpur那尔补罗() feud.Barony {
	return c.那尔补罗Lalpur
}

var CDvaraka堕罗迦 DvarakaCounty = &堕罗迦DvarakaCounty{}

func init() {
	f := CDvaraka堕罗迦.(*堕罗迦DvarakaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1136",
		Title:     "dvaraka",
		TitleName: "堕罗迦",
		TitleCode: "c_dvaraka",
		Baronies:  map[string]feud.Barony{},
	}

	f.般跋Bhanvad = BBhanvad般跋
	f.般跋Bhanvad.SetParent(f)

	f.堕罗迦Dvaraka = BDvaraka堕罗迦
	f.堕罗迦Dvaraka.SetParent(f)

	f.堕罗迦地舍Dwarakadheesh = BDwarakadheesh堕罗迦地舍
	f.堕罗迦地舍Dwarakadheesh.SetParent(f)

	f.戈摩蒂Gomati = BGomati戈摩蒂
	f.戈摩蒂Gomati.SetParent(f)

	f.呼旬提Huljanti = BHuljanti呼旬提
	f.呼旬提Huljanti.SetParent(f)

	f.阇羯罗Jakar = BJakar阇羯罗
	f.阇羯罗Jakar.SetParent(f)

	f.那尔补罗Lalpur = BLalpur那尔补罗
	f.那尔补罗Lalpur.SetParent(f)

}
