package bihar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BiharCounty interface {
    feud.County
    BBihar比豪尔() 	feud.Barony
    BBiharkeresztes比豪尔凯赖斯泰什() 	feud.Barony
    BDebrecen德布勒森() 	feud.Barony
    BElesd埃莱什特() 	feud.Barony
    BNagybajom瑙吉鲍约姆() 	feud.Barony
    BNagyvarad瑙吉瓦劳德() 	feud.Barony
    BSzalard萨拉尔德() 	feud.Barony
    BZolonta佐龙陶() 	feud.Barony
}

type 比豪尔BiharCounty struct {
	feud.BaseCounty
	比豪尔Bihar 	feud.Barony
	比豪尔凯赖斯泰什Biharkeresztes 	feud.Barony
	德布勒森Debrecen 	feud.Barony
	埃莱什特Elesd 	feud.Barony
	瑙吉鲍约姆Nagybajom 	feud.Barony
	瑙吉瓦劳德Nagyvarad 	feud.Barony
	萨拉尔德Szalard 	feud.Barony
	佐龙陶Zolonta 	feud.Barony
}

func (c *比豪尔BiharCounty) BBihar比豪尔() feud.Barony {
	return c.比豪尔Bihar
}
    
func (c *比豪尔BiharCounty) BBiharkeresztes比豪尔凯赖斯泰什() feud.Barony {
	return c.比豪尔凯赖斯泰什Biharkeresztes
}
    
func (c *比豪尔BiharCounty) BDebrecen德布勒森() feud.Barony {
	return c.德布勒森Debrecen
}
    
func (c *比豪尔BiharCounty) BElesd埃莱什特() feud.Barony {
	return c.埃莱什特Elesd
}
    
func (c *比豪尔BiharCounty) BNagybajom瑙吉鲍约姆() feud.Barony {
	return c.瑙吉鲍约姆Nagybajom
}
    
func (c *比豪尔BiharCounty) BNagyvarad瑙吉瓦劳德() feud.Barony {
	return c.瑙吉瓦劳德Nagyvarad
}
    
func (c *比豪尔BiharCounty) BSzalard萨拉尔德() feud.Barony {
	return c.萨拉尔德Szalard
}
    
func (c *比豪尔BiharCounty) BZolonta佐龙陶() feud.Barony {
	return c.佐龙陶Zolonta
}
    
var CBihar比豪尔 BiharCounty = &比豪尔BiharCounty{}

func init() {
	f := CBihar比豪尔.(*比豪尔BiharCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "520",
		Title:     "bihar",
		TitleName: "比豪尔",
		TitleCode: "c_bihar",
		Baronies:  map[string]feud.Barony{},
	}

	f.比豪尔Bihar = BBihar比豪尔
	f.比豪尔Bihar.SetParent(f)
	
	f.比豪尔凯赖斯泰什Biharkeresztes = BBiharkeresztes比豪尔凯赖斯泰什
	f.比豪尔凯赖斯泰什Biharkeresztes.SetParent(f)
	
	f.德布勒森Debrecen = BDebrecen德布勒森
	f.德布勒森Debrecen.SetParent(f)
	
	f.埃莱什特Elesd = BElesd埃莱什特
	f.埃莱什特Elesd.SetParent(f)
	
	f.瑙吉鲍约姆Nagybajom = BNagybajom瑙吉鲍约姆
	f.瑙吉鲍约姆Nagybajom.SetParent(f)
	
	f.瑙吉瓦劳德Nagyvarad = BNagyvarad瑙吉瓦劳德
	f.瑙吉瓦劳德Nagyvarad.SetParent(f)
	
	f.萨拉尔德Szalard = BSzalard萨拉尔德
	f.萨拉尔德Szalard.SetParent(f)
	
	f.佐龙陶Zolonta = BZolonta佐龙陶
	f.佐龙陶Zolonta.SetParent(f)
	
}
