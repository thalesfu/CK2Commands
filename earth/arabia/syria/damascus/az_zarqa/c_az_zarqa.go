package az_zarqa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Az_zarqaCounty interface {
    feud.County
    BAmratamad阿木拉特玛德() 	feud.Barony
    BHashemiyya哈希姆() 	feud.Barony
    BKhurqah乌姆胡尔盖() 	feud.Barony
    BQasralhallabat哈拉巴特堡() 	feud.Barony
    BQasramra阿姆拉堡() 	feud.Barony
    BRusseifa鲁赛法() 	feud.Barony
    BShabib沙比布() 	feud.Barony
    BZarqa扎尔卡() 	feud.Barony
}

type 扎尔卡Az_zarqaCounty struct {
	feud.BaseCounty
	阿木拉特玛德Amratamad 	feud.Barony
	哈希姆Hashemiyya 	feud.Barony
	乌姆胡尔盖Khurqah 	feud.Barony
	哈拉巴特堡Qasralhallabat 	feud.Barony
	阿姆拉堡Qasramra 	feud.Barony
	鲁赛法Russeifa 	feud.Barony
	沙比布Shabib 	feud.Barony
	扎尔卡Zarqa 	feud.Barony
}

func (c *扎尔卡Az_zarqaCounty) BAmratamad阿木拉特玛德() feud.Barony {
	return c.阿木拉特玛德Amratamad
}
    
func (c *扎尔卡Az_zarqaCounty) BHashemiyya哈希姆() feud.Barony {
	return c.哈希姆Hashemiyya
}
    
func (c *扎尔卡Az_zarqaCounty) BKhurqah乌姆胡尔盖() feud.Barony {
	return c.乌姆胡尔盖Khurqah
}
    
func (c *扎尔卡Az_zarqaCounty) BQasralhallabat哈拉巴特堡() feud.Barony {
	return c.哈拉巴特堡Qasralhallabat
}
    
func (c *扎尔卡Az_zarqaCounty) BQasramra阿姆拉堡() feud.Barony {
	return c.阿姆拉堡Qasramra
}
    
func (c *扎尔卡Az_zarqaCounty) BRusseifa鲁赛法() feud.Barony {
	return c.鲁赛法Russeifa
}
    
func (c *扎尔卡Az_zarqaCounty) BShabib沙比布() feud.Barony {
	return c.沙比布Shabib
}
    
func (c *扎尔卡Az_zarqaCounty) BZarqa扎尔卡() feud.Barony {
	return c.扎尔卡Zarqa
}
    
var CAz_zarqa扎尔卡 Az_zarqaCounty = &扎尔卡Az_zarqaCounty{}

func init() {
	f := CAz_zarqa扎尔卡.(*扎尔卡Az_zarqaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "714",
		Title:     "az_zarqa",
		TitleName: "扎尔卡",
		TitleCode: "c_az_zarqa",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿木拉特玛德Amratamad = BAmratamad阿木拉特玛德
	f.阿木拉特玛德Amratamad.SetParent(f)
	
	f.哈希姆Hashemiyya = BHashemiyya哈希姆
	f.哈希姆Hashemiyya.SetParent(f)
	
	f.乌姆胡尔盖Khurqah = BKhurqah乌姆胡尔盖
	f.乌姆胡尔盖Khurqah.SetParent(f)
	
	f.哈拉巴特堡Qasralhallabat = BQasralhallabat哈拉巴特堡
	f.哈拉巴特堡Qasralhallabat.SetParent(f)
	
	f.阿姆拉堡Qasramra = BQasramra阿姆拉堡
	f.阿姆拉堡Qasramra.SetParent(f)
	
	f.鲁赛法Russeifa = BRusseifa鲁赛法
	f.鲁赛法Russeifa.SetParent(f)
	
	f.沙比布Shabib = BShabib沙比布
	f.沙比布Shabib.SetParent(f)
	
	f.扎尔卡Zarqa = BZarqa扎尔卡
	f.扎尔卡Zarqa.SetParent(f)
	
}
