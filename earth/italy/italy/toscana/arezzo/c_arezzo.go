package arezzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArezzoCounty interface {
    feud.County
    BArezzo阿雷佐() 	feud.Barony
    BCastiglione_arezzo卡斯蒂廖内() 	feud.Barony
    BCivitella_del_vescovo奇维泰拉韦斯科沃() 	feud.Barony
    BCortona科尔托纳() 	feud.Barony
    BLucignano卢奇尼亚诺() 	feud.Barony
    BMontevarchi蒙特瓦尔基() 	feud.Barony
    BSansepolcro圣塞波尔克罗() 	feud.Barony
}

type 阿雷佐ArezzoCounty struct {
	feud.BaseCounty
	阿雷佐Arezzo 	feud.Barony
	卡斯蒂廖内Castiglione_arezzo 	feud.Barony
	奇维泰拉韦斯科沃Civitella_del_vescovo 	feud.Barony
	科尔托纳Cortona 	feud.Barony
	卢奇尼亚诺Lucignano 	feud.Barony
	蒙特瓦尔基Montevarchi 	feud.Barony
	圣塞波尔克罗Sansepolcro 	feud.Barony
}

func (c *阿雷佐ArezzoCounty) BArezzo阿雷佐() feud.Barony {
	return c.阿雷佐Arezzo
}
    
func (c *阿雷佐ArezzoCounty) BCastiglione_arezzo卡斯蒂廖内() feud.Barony {
	return c.卡斯蒂廖内Castiglione_arezzo
}
    
func (c *阿雷佐ArezzoCounty) BCivitella_del_vescovo奇维泰拉韦斯科沃() feud.Barony {
	return c.奇维泰拉韦斯科沃Civitella_del_vescovo
}
    
func (c *阿雷佐ArezzoCounty) BCortona科尔托纳() feud.Barony {
	return c.科尔托纳Cortona
}
    
func (c *阿雷佐ArezzoCounty) BLucignano卢奇尼亚诺() feud.Barony {
	return c.卢奇尼亚诺Lucignano
}
    
func (c *阿雷佐ArezzoCounty) BMontevarchi蒙特瓦尔基() feud.Barony {
	return c.蒙特瓦尔基Montevarchi
}
    
func (c *阿雷佐ArezzoCounty) BSansepolcro圣塞波尔克罗() feud.Barony {
	return c.圣塞波尔克罗Sansepolcro
}
    
var CArezzo阿雷佐 ArezzoCounty = &阿雷佐ArezzoCounty{}

func init() {
	f := CArezzo阿雷佐.(*阿雷佐ArezzoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1713",
		Title:     "arezzo",
		TitleName: "阿雷佐",
		TitleCode: "c_arezzo",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿雷佐Arezzo = BArezzo阿雷佐
	f.阿雷佐Arezzo.SetParent(f)
	
	f.卡斯蒂廖内Castiglione_arezzo = BCastiglione_arezzo卡斯蒂廖内
	f.卡斯蒂廖内Castiglione_arezzo.SetParent(f)
	
	f.奇维泰拉韦斯科沃Civitella_del_vescovo = BCivitella_del_vescovo奇维泰拉韦斯科沃
	f.奇维泰拉韦斯科沃Civitella_del_vescovo.SetParent(f)
	
	f.科尔托纳Cortona = BCortona科尔托纳
	f.科尔托纳Cortona.SetParent(f)
	
	f.卢奇尼亚诺Lucignano = BLucignano卢奇尼亚诺
	f.卢奇尼亚诺Lucignano.SetParent(f)
	
	f.蒙特瓦尔基Montevarchi = BMontevarchi蒙特瓦尔基
	f.蒙特瓦尔基Montevarchi.SetParent(f)
	
	f.圣塞波尔克罗Sansepolcro = BSansepolcro圣塞波尔克罗
	f.圣塞波尔克罗Sansepolcro.SetParent(f)
	
}
