package baalbek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BaalbekCounty interface {
    feud.County
    BAkkar阿卡尔() 	feud.Barony
    BBaalbek巴勒贝克() 	feud.Barony
    BBuissera布瑟塞拉() 	feud.Barony
    BChlifa什里法() 	feud.Barony
    BHalbah哈勒巴() 	feud.Barony
    BKrakdeschevaliers骑士堡() 	feud.Barony
    BLaboue莱卜沃() 	feud.Barony
    BRiyaq里亚格() 	feud.Barony
    BZahle扎赫勒() 	feud.Barony
}

type 巴勒贝克BaalbekCounty struct {
	feud.BaseCounty
	阿卡尔Akkar 	feud.Barony
	巴勒贝克Baalbek 	feud.Barony
	布瑟塞拉Buissera 	feud.Barony
	什里法Chlifa 	feud.Barony
	哈勒巴Halbah 	feud.Barony
	骑士堡Krakdeschevaliers 	feud.Barony
	莱卜沃Laboue 	feud.Barony
	里亚格Riyaq 	feud.Barony
	扎赫勒Zahle 	feud.Barony
}

func (c *巴勒贝克BaalbekCounty) BAkkar阿卡尔() feud.Barony {
	return c.阿卡尔Akkar
}
    
func (c *巴勒贝克BaalbekCounty) BBaalbek巴勒贝克() feud.Barony {
	return c.巴勒贝克Baalbek
}
    
func (c *巴勒贝克BaalbekCounty) BBuissera布瑟塞拉() feud.Barony {
	return c.布瑟塞拉Buissera
}
    
func (c *巴勒贝克BaalbekCounty) BChlifa什里法() feud.Barony {
	return c.什里法Chlifa
}
    
func (c *巴勒贝克BaalbekCounty) BHalbah哈勒巴() feud.Barony {
	return c.哈勒巴Halbah
}
    
func (c *巴勒贝克BaalbekCounty) BKrakdeschevaliers骑士堡() feud.Barony {
	return c.骑士堡Krakdeschevaliers
}
    
func (c *巴勒贝克BaalbekCounty) BLaboue莱卜沃() feud.Barony {
	return c.莱卜沃Laboue
}
    
func (c *巴勒贝克BaalbekCounty) BRiyaq里亚格() feud.Barony {
	return c.里亚格Riyaq
}
    
func (c *巴勒贝克BaalbekCounty) BZahle扎赫勒() feud.Barony {
	return c.扎赫勒Zahle
}
    
var CBaalbek巴勒贝克 BaalbekCounty = &巴勒贝克BaalbekCounty{}

func init() {
	f := CBaalbek巴勒贝克.(*巴勒贝克BaalbekCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "768",
		Title:     "baalbek",
		TitleName: "巴勒贝克",
		TitleCode: "c_baalbek",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卡尔Akkar = BAkkar阿卡尔
	f.阿卡尔Akkar.SetParent(f)
	
	f.巴勒贝克Baalbek = BBaalbek巴勒贝克
	f.巴勒贝克Baalbek.SetParent(f)
	
	f.布瑟塞拉Buissera = BBuissera布瑟塞拉
	f.布瑟塞拉Buissera.SetParent(f)
	
	f.什里法Chlifa = BChlifa什里法
	f.什里法Chlifa.SetParent(f)
	
	f.哈勒巴Halbah = BHalbah哈勒巴
	f.哈勒巴Halbah.SetParent(f)
	
	f.骑士堡Krakdeschevaliers = BKrakdeschevaliers骑士堡
	f.骑士堡Krakdeschevaliers.SetParent(f)
	
	f.莱卜沃Laboue = BLaboue莱卜沃
	f.莱卜沃Laboue.SetParent(f)
	
	f.里亚格Riyaq = BRiyaq里亚格
	f.里亚格Riyaq.SetParent(f)
	
	f.扎赫勒Zahle = BZahle扎赫勒
	f.扎赫勒Zahle.SetParent(f)
	
}
