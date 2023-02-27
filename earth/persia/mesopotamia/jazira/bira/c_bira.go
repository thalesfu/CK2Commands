package bira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BiraCounty interface {
    feud.County
    BBira比雷() 	feud.Barony
    BDerik代里克() 	feud.Barony
    BQalarebete喀拉雷比特() 	feud.Barony
    BQoser喀瑟尔() 	feud.Barony
    BResaina拉斯艾因() 	feud.Barony
    BSamrah萨梅拉赫() 	feud.Barony
    BStewr斯特维() 	feud.Barony
    BTella特拉拉() 	feud.Barony
}

type 比雷BiraCounty struct {
	feud.BaseCounty
	比雷Bira 	feud.Barony
	代里克Derik 	feud.Barony
	喀拉雷比特Qalarebete 	feud.Barony
	喀瑟尔Qoser 	feud.Barony
	拉斯艾因Resaina 	feud.Barony
	萨梅拉赫Samrah 	feud.Barony
	斯特维Stewr 	feud.Barony
	特拉拉Tella 	feud.Barony
}

func (c *比雷BiraCounty) BBira比雷() feud.Barony {
	return c.比雷Bira
}
    
func (c *比雷BiraCounty) BDerik代里克() feud.Barony {
	return c.代里克Derik
}
    
func (c *比雷BiraCounty) BQalarebete喀拉雷比特() feud.Barony {
	return c.喀拉雷比特Qalarebete
}
    
func (c *比雷BiraCounty) BQoser喀瑟尔() feud.Barony {
	return c.喀瑟尔Qoser
}
    
func (c *比雷BiraCounty) BResaina拉斯艾因() feud.Barony {
	return c.拉斯艾因Resaina
}
    
func (c *比雷BiraCounty) BSamrah萨梅拉赫() feud.Barony {
	return c.萨梅拉赫Samrah
}
    
func (c *比雷BiraCounty) BStewr斯特维() feud.Barony {
	return c.斯特维Stewr
}
    
func (c *比雷BiraCounty) BTella特拉拉() feud.Barony {
	return c.特拉拉Tella
}
    
var CBira比雷 BiraCounty = &比雷BiraCounty{}

func init() {
	f := CBira比雷.(*比雷BiraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "700",
		Title:     "bira",
		TitleName: "比雷",
		TitleCode: "c_bira",
		Baronies:  map[string]feud.Barony{},
	}

	f.比雷Bira = BBira比雷
	f.比雷Bira.SetParent(f)
	
	f.代里克Derik = BDerik代里克
	f.代里克Derik.SetParent(f)
	
	f.喀拉雷比特Qalarebete = BQalarebete喀拉雷比特
	f.喀拉雷比特Qalarebete.SetParent(f)
	
	f.喀瑟尔Qoser = BQoser喀瑟尔
	f.喀瑟尔Qoser.SetParent(f)
	
	f.拉斯艾因Resaina = BResaina拉斯艾因
	f.拉斯艾因Resaina.SetParent(f)
	
	f.萨梅拉赫Samrah = BSamrah萨梅拉赫
	f.萨梅拉赫Samrah.SetParent(f)
	
	f.斯特维Stewr = BStewr斯特维
	f.斯特维Stewr.SetParent(f)
	
	f.特拉拉Tella = BTella特拉拉
	f.特拉拉Tella.SetParent(f)
	
}
