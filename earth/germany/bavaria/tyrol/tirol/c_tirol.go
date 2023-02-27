package tirol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TirolCounty interface {
    feud.County
    BBregenz布雷根茨() 	feud.Barony
    BDornbirn多恩比恩() 	feud.Barony
    BImst伊姆斯特() 	feud.Barony
    BLandeck兰德克() 	feud.Barony
    BNenzing嫩青() 	feud.Barony
    BRied里德() 	feud.Barony
    BStanton圣安东() 	feud.Barony
    BSterzing施泰尔青() 	feud.Barony
}

type 兰德克TirolCounty struct {
	feud.BaseCounty
	布雷根茨Bregenz 	feud.Barony
	多恩比恩Dornbirn 	feud.Barony
	伊姆斯特Imst 	feud.Barony
	兰德克Landeck 	feud.Barony
	嫩青Nenzing 	feud.Barony
	里德Ried 	feud.Barony
	圣安东Stanton 	feud.Barony
	施泰尔青Sterzing 	feud.Barony
}

func (c *兰德克TirolCounty) BBregenz布雷根茨() feud.Barony {
	return c.布雷根茨Bregenz
}
    
func (c *兰德克TirolCounty) BDornbirn多恩比恩() feud.Barony {
	return c.多恩比恩Dornbirn
}
    
func (c *兰德克TirolCounty) BImst伊姆斯特() feud.Barony {
	return c.伊姆斯特Imst
}
    
func (c *兰德克TirolCounty) BLandeck兰德克() feud.Barony {
	return c.兰德克Landeck
}
    
func (c *兰德克TirolCounty) BNenzing嫩青() feud.Barony {
	return c.嫩青Nenzing
}
    
func (c *兰德克TirolCounty) BRied里德() feud.Barony {
	return c.里德Ried
}
    
func (c *兰德克TirolCounty) BStanton圣安东() feud.Barony {
	return c.圣安东Stanton
}
    
func (c *兰德克TirolCounty) BSterzing施泰尔青() feud.Barony {
	return c.施泰尔青Sterzing
}
    
var CTirol兰德克 TirolCounty = &兰德克TirolCounty{}

func init() {
	f := CTirol兰德克.(*兰德克TirolCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "316",
		Title:     "tirol",
		TitleName: "兰德克",
		TitleCode: "c_tirol",
		Baronies:  map[string]feud.Barony{},
	}

	f.布雷根茨Bregenz = BBregenz布雷根茨
	f.布雷根茨Bregenz.SetParent(f)
	
	f.多恩比恩Dornbirn = BDornbirn多恩比恩
	f.多恩比恩Dornbirn.SetParent(f)
	
	f.伊姆斯特Imst = BImst伊姆斯特
	f.伊姆斯特Imst.SetParent(f)
	
	f.兰德克Landeck = BLandeck兰德克
	f.兰德克Landeck.SetParent(f)
	
	f.嫩青Nenzing = BNenzing嫩青
	f.嫩青Nenzing.SetParent(f)
	
	f.里德Ried = BRied里德
	f.里德Ried.SetParent(f)
	
	f.圣安东Stanton = BStanton圣安东
	f.圣安东Stanton.SetParent(f)
	
	f.施泰尔青Sterzing = BSterzing施泰尔青
	f.施泰尔青Sterzing.SetParent(f)
	
}
