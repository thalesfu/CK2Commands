package litomerice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LitomericeCounty interface {
    feud.County
    BBoleslav博莱斯拉夫() 	feud.Barony
    BCeskalipa捷克利帕() 	feud.Barony
    BDecin杰钦() 	feud.Barony
    BDuba_litomerice杜巴() 	feud.Barony
    BLitomerice利托梅日采() 	feud.Barony
    BMelnik_bohemia梅尔尼克() 	feud.Barony
    BMnichovo_hradiste慕尼黑城堡() 	feud.Barony
    BNymburk宁布尔克() 	feud.Barony
    BZitava日塔瓦() 	feud.Barony
}

type 利托梅日采LitomericeCounty struct {
	feud.BaseCounty
	博莱斯拉夫Boleslav 	feud.Barony
	捷克利帕Ceskalipa 	feud.Barony
	杰钦Decin 	feud.Barony
	杜巴Duba_litomerice 	feud.Barony
	利托梅日采Litomerice 	feud.Barony
	梅尔尼克Melnik_bohemia 	feud.Barony
	慕尼黑城堡Mnichovo_hradiste 	feud.Barony
	宁布尔克Nymburk 	feud.Barony
	日塔瓦Zitava 	feud.Barony
}

func (c *利托梅日采LitomericeCounty) BBoleslav博莱斯拉夫() feud.Barony {
	return c.博莱斯拉夫Boleslav
}
    
func (c *利托梅日采LitomericeCounty) BCeskalipa捷克利帕() feud.Barony {
	return c.捷克利帕Ceskalipa
}
    
func (c *利托梅日采LitomericeCounty) BDecin杰钦() feud.Barony {
	return c.杰钦Decin
}
    
func (c *利托梅日采LitomericeCounty) BDuba_litomerice杜巴() feud.Barony {
	return c.杜巴Duba_litomerice
}
    
func (c *利托梅日采LitomericeCounty) BLitomerice利托梅日采() feud.Barony {
	return c.利托梅日采Litomerice
}
    
func (c *利托梅日采LitomericeCounty) BMelnik_bohemia梅尔尼克() feud.Barony {
	return c.梅尔尼克Melnik_bohemia
}
    
func (c *利托梅日采LitomericeCounty) BMnichovo_hradiste慕尼黑城堡() feud.Barony {
	return c.慕尼黑城堡Mnichovo_hradiste
}
    
func (c *利托梅日采LitomericeCounty) BNymburk宁布尔克() feud.Barony {
	return c.宁布尔克Nymburk
}
    
func (c *利托梅日采LitomericeCounty) BZitava日塔瓦() feud.Barony {
	return c.日塔瓦Zitava
}
    
var CLitomerice利托梅日采 LitomericeCounty = &利托梅日采LitomericeCounty{}

func init() {
	f := CLitomerice利托梅日采.(*利托梅日采LitomericeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "363",
		Title:     "litomerice",
		TitleName: "利托梅日采",
		TitleCode: "c_litomerice",
		Baronies:  map[string]feud.Barony{},
	}

	f.博莱斯拉夫Boleslav = BBoleslav博莱斯拉夫
	f.博莱斯拉夫Boleslav.SetParent(f)
	
	f.捷克利帕Ceskalipa = BCeskalipa捷克利帕
	f.捷克利帕Ceskalipa.SetParent(f)
	
	f.杰钦Decin = BDecin杰钦
	f.杰钦Decin.SetParent(f)
	
	f.杜巴Duba_litomerice = BDuba_litomerice杜巴
	f.杜巴Duba_litomerice.SetParent(f)
	
	f.利托梅日采Litomerice = BLitomerice利托梅日采
	f.利托梅日采Litomerice.SetParent(f)
	
	f.梅尔尼克Melnik_bohemia = BMelnik_bohemia梅尔尼克
	f.梅尔尼克Melnik_bohemia.SetParent(f)
	
	f.慕尼黑城堡Mnichovo_hradiste = BMnichovo_hradiste慕尼黑城堡
	f.慕尼黑城堡Mnichovo_hradiste.SetParent(f)
	
	f.宁布尔克Nymburk = BNymburk宁布尔克
	f.宁布尔克Nymburk.SetParent(f)
	
	f.日塔瓦Zitava = BZitava日塔瓦
	f.日塔瓦Zitava.SetParent(f)
	
}
