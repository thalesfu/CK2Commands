package naimisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NaimisaCounty interface {
	feud.County
	BBahraich婆诃罗伊支() feud.Barony
	BDurbate补婆谛() feud.Barony
	BGonda郡荼() feud.Barony
	BLakhimpur罗企摩城() feud.Barony
	BMisrikh弥室利迦() feud.Barony
	BNaimisa泥弥沙() feud.Barony
	BSitapur私多城() feud.Barony
}

type 泥弥沙NaimisaCounty struct {
	feud.BaseCounty
	婆诃罗伊支Bahraich feud.Barony
	补婆谛Durbate    feud.Barony
	郡荼Gonda       feud.Barony
	罗企摩城Lakhimpur feud.Barony
	弥室利迦Misrikh   feud.Barony
	泥弥沙Naimisa    feud.Barony
	私多城Sitapur    feud.Barony
}

func (c *泥弥沙NaimisaCounty) BBahraich婆诃罗伊支() feud.Barony {
	return c.婆诃罗伊支Bahraich
}

func (c *泥弥沙NaimisaCounty) BDurbate补婆谛() feud.Barony {
	return c.补婆谛Durbate
}

func (c *泥弥沙NaimisaCounty) BGonda郡荼() feud.Barony {
	return c.郡荼Gonda
}

func (c *泥弥沙NaimisaCounty) BLakhimpur罗企摩城() feud.Barony {
	return c.罗企摩城Lakhimpur
}

func (c *泥弥沙NaimisaCounty) BMisrikh弥室利迦() feud.Barony {
	return c.弥室利迦Misrikh
}

func (c *泥弥沙NaimisaCounty) BNaimisa泥弥沙() feud.Barony {
	return c.泥弥沙Naimisa
}

func (c *泥弥沙NaimisaCounty) BSitapur私多城() feud.Barony {
	return c.私多城Sitapur
}

var CNaimisa泥弥沙 NaimisaCounty = &泥弥沙NaimisaCounty{}

func init() {
	f := CNaimisa泥弥沙.(*泥弥沙NaimisaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1422",
		Title:     "naimisa",
		TitleName: "泥弥沙",
		TitleCode: "c_naimisa",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆诃罗伊支Bahraich = BBahraich婆诃罗伊支
	f.婆诃罗伊支Bahraich.SetParent(f)

	f.补婆谛Durbate = BDurbate补婆谛
	f.补婆谛Durbate.SetParent(f)

	f.郡荼Gonda = BGonda郡荼
	f.郡荼Gonda.SetParent(f)

	f.罗企摩城Lakhimpur = BLakhimpur罗企摩城
	f.罗企摩城Lakhimpur.SetParent(f)

	f.弥室利迦Misrikh = BMisrikh弥室利迦
	f.弥室利迦Misrikh.SetParent(f)

	f.泥弥沙Naimisa = BNaimisa泥弥沙
	f.泥弥沙Naimisa.SetParent(f)

	f.私多城Sitapur = BSitapur私多城
	f.私多城Sitapur.SetParent(f)

}
