package chitrakut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChitrakutCounty interface {
	feud.County
	BBaroli婆卢利() feud.Barony
	BBhainsrorgarh拜恩斯罗加尔() feud.Barony
	BChitrakut质多罗矩吒() feud.Barony
	BMandalgarh曼荼罗姞利呬() feud.Barony
	BNagari那伽梨() feud.Barony
	BRampura罗摩补罗() feud.Barony
	BSaunk绍恩克() feud.Barony
}

type 质多罗矩吒ChitrakutCounty struct {
	feud.BaseCounty
	婆卢利Baroli           feud.Barony
	拜恩斯罗加尔Bhainsrorgarh feud.Barony
	质多罗矩吒Chitrakut      feud.Barony
	曼荼罗姞利呬Mandalgarh    feud.Barony
	那伽梨Nagari           feud.Barony
	罗摩补罗Rampura         feud.Barony
	绍恩克Saunk            feud.Barony
}

func (c *质多罗矩吒ChitrakutCounty) BBaroli婆卢利() feud.Barony {
	return c.婆卢利Baroli
}

func (c *质多罗矩吒ChitrakutCounty) BBhainsrorgarh拜恩斯罗加尔() feud.Barony {
	return c.拜恩斯罗加尔Bhainsrorgarh
}

func (c *质多罗矩吒ChitrakutCounty) BChitrakut质多罗矩吒() feud.Barony {
	return c.质多罗矩吒Chitrakut
}

func (c *质多罗矩吒ChitrakutCounty) BMandalgarh曼荼罗姞利呬() feud.Barony {
	return c.曼荼罗姞利呬Mandalgarh
}

func (c *质多罗矩吒ChitrakutCounty) BNagari那伽梨() feud.Barony {
	return c.那伽梨Nagari
}

func (c *质多罗矩吒ChitrakutCounty) BRampura罗摩补罗() feud.Barony {
	return c.罗摩补罗Rampura
}

func (c *质多罗矩吒ChitrakutCounty) BSaunk绍恩克() feud.Barony {
	return c.绍恩克Saunk
}

var CChitrakut质多罗矩吒 ChitrakutCounty = &质多罗矩吒ChitrakutCounty{}

func init() {
	f := CChitrakut质多罗矩吒.(*质多罗矩吒ChitrakutCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1302",
		Title:     "chitrakut",
		TitleName: "质多罗矩吒",
		TitleCode: "c_chitrakut",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆卢利Baroli = BBaroli婆卢利
	f.婆卢利Baroli.SetParent(f)

	f.拜恩斯罗加尔Bhainsrorgarh = BBhainsrorgarh拜恩斯罗加尔
	f.拜恩斯罗加尔Bhainsrorgarh.SetParent(f)

	f.质多罗矩吒Chitrakut = BChitrakut质多罗矩吒
	f.质多罗矩吒Chitrakut.SetParent(f)

	f.曼荼罗姞利呬Mandalgarh = BMandalgarh曼荼罗姞利呬
	f.曼荼罗姞利呬Mandalgarh.SetParent(f)

	f.那伽梨Nagari = BNagari那伽梨
	f.那伽梨Nagari.SetParent(f)

	f.罗摩补罗Rampura = BRampura罗摩补罗
	f.罗摩补罗Rampura.SetParent(f)

	f.绍恩克Saunk = BSaunk绍恩克
	f.绍恩克Saunk.SetParent(f)

}
