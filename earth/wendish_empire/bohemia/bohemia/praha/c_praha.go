package praha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PrahaCounty interface {
    feud.County
    BBrevnov布热夫诺夫() 	feud.Barony
    BKolin科林() 	feud.Barony
    BKourim考日姆() 	feud.Barony
    BKuttenberg库滕贝格() 	feud.Barony
    BPraha布拉格() 	feud.Barony
    BSazava萨扎瓦() 	feud.Barony
    BStare_mesto斯塔雷梅斯托() 	feud.Barony
    BVysehrad高堡() 	feud.Barony
    BZbraslav兹布拉斯拉夫() 	feud.Barony
}

type 布拉格PrahaCounty struct {
	feud.BaseCounty
	布热夫诺夫Brevnov 	feud.Barony
	科林Kolin 	feud.Barony
	考日姆Kourim 	feud.Barony
	库滕贝格Kuttenberg 	feud.Barony
	布拉格Praha 	feud.Barony
	萨扎瓦Sazava 	feud.Barony
	斯塔雷梅斯托Stare_mesto 	feud.Barony
	高堡Vysehrad 	feud.Barony
	兹布拉斯拉夫Zbraslav 	feud.Barony
}

func (c *布拉格PrahaCounty) BBrevnov布热夫诺夫() feud.Barony {
	return c.布热夫诺夫Brevnov
}
    
func (c *布拉格PrahaCounty) BKolin科林() feud.Barony {
	return c.科林Kolin
}
    
func (c *布拉格PrahaCounty) BKourim考日姆() feud.Barony {
	return c.考日姆Kourim
}
    
func (c *布拉格PrahaCounty) BKuttenberg库滕贝格() feud.Barony {
	return c.库滕贝格Kuttenberg
}
    
func (c *布拉格PrahaCounty) BPraha布拉格() feud.Barony {
	return c.布拉格Praha
}
    
func (c *布拉格PrahaCounty) BSazava萨扎瓦() feud.Barony {
	return c.萨扎瓦Sazava
}
    
func (c *布拉格PrahaCounty) BStare_mesto斯塔雷梅斯托() feud.Barony {
	return c.斯塔雷梅斯托Stare_mesto
}
    
func (c *布拉格PrahaCounty) BVysehrad高堡() feud.Barony {
	return c.高堡Vysehrad
}
    
func (c *布拉格PrahaCounty) BZbraslav兹布拉斯拉夫() feud.Barony {
	return c.兹布拉斯拉夫Zbraslav
}
    
var CPraha布拉格 PrahaCounty = &布拉格PrahaCounty{}

func init() {
	f := CPraha布拉格.(*布拉格PrahaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "437",
		Title:     "praha",
		TitleName: "布拉格",
		TitleCode: "c_praha",
		Baronies:  map[string]feud.Barony{},
	}

	f.布热夫诺夫Brevnov = BBrevnov布热夫诺夫
	f.布热夫诺夫Brevnov.SetParent(f)
	
	f.科林Kolin = BKolin科林
	f.科林Kolin.SetParent(f)
	
	f.考日姆Kourim = BKourim考日姆
	f.考日姆Kourim.SetParent(f)
	
	f.库滕贝格Kuttenberg = BKuttenberg库滕贝格
	f.库滕贝格Kuttenberg.SetParent(f)
	
	f.布拉格Praha = BPraha布拉格
	f.布拉格Praha.SetParent(f)
	
	f.萨扎瓦Sazava = BSazava萨扎瓦
	f.萨扎瓦Sazava.SetParent(f)
	
	f.斯塔雷梅斯托Stare_mesto = BStare_mesto斯塔雷梅斯托
	f.斯塔雷梅斯托Stare_mesto.SetParent(f)
	
	f.高堡Vysehrad = BVysehrad高堡
	f.高堡Vysehrad.SetParent(f)
	
	f.兹布拉斯拉夫Zbraslav = BZbraslav兹布拉斯拉夫
	f.兹布拉斯拉夫Zbraslav.SetParent(f)
	
}
