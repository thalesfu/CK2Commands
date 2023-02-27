package kharga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhargaCounty interface {
    feud.County
    BBanja班贾() 	feud.Barony
    BBaris巴里斯() 	feud.Barony
    BElkharga哈里杰() 	feud.Barony
    BGahdam杰赫达姆() 	feud.Barony
    BIsmant艾斯曼() 	feud.Barony
    BPanopolis帕诺波利斯() 	feud.Barony
    BZuhrah祖赫拉() 	feud.Barony
}

type 哈里杰KhargaCounty struct {
	feud.BaseCounty
	班贾Banja 	feud.Barony
	巴里斯Baris 	feud.Barony
	哈里杰Elkharga 	feud.Barony
	杰赫达姆Gahdam 	feud.Barony
	艾斯曼Ismant 	feud.Barony
	帕诺波利斯Panopolis 	feud.Barony
	祖赫拉Zuhrah 	feud.Barony
}

func (c *哈里杰KhargaCounty) BBanja班贾() feud.Barony {
	return c.班贾Banja
}
    
func (c *哈里杰KhargaCounty) BBaris巴里斯() feud.Barony {
	return c.巴里斯Baris
}
    
func (c *哈里杰KhargaCounty) BElkharga哈里杰() feud.Barony {
	return c.哈里杰Elkharga
}
    
func (c *哈里杰KhargaCounty) BGahdam杰赫达姆() feud.Barony {
	return c.杰赫达姆Gahdam
}
    
func (c *哈里杰KhargaCounty) BIsmant艾斯曼() feud.Barony {
	return c.艾斯曼Ismant
}
    
func (c *哈里杰KhargaCounty) BPanopolis帕诺波利斯() feud.Barony {
	return c.帕诺波利斯Panopolis
}
    
func (c *哈里杰KhargaCounty) BZuhrah祖赫拉() feud.Barony {
	return c.祖赫拉Zuhrah
}
    
var CKharga哈里杰 KhargaCounty = &哈里杰KhargaCounty{}

func init() {
	f := CKharga哈里杰.(*哈里杰KhargaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1782",
		Title:     "kharga",
		TitleName: "哈里杰",
		TitleCode: "c_kharga",
		Baronies:  map[string]feud.Barony{},
	}

	f.班贾Banja = BBanja班贾
	f.班贾Banja.SetParent(f)
	
	f.巴里斯Baris = BBaris巴里斯
	f.巴里斯Baris.SetParent(f)
	
	f.哈里杰Elkharga = BElkharga哈里杰
	f.哈里杰Elkharga.SetParent(f)
	
	f.杰赫达姆Gahdam = BGahdam杰赫达姆
	f.杰赫达姆Gahdam.SetParent(f)
	
	f.艾斯曼Ismant = BIsmant艾斯曼
	f.艾斯曼Ismant.SetParent(f)
	
	f.帕诺波利斯Panopolis = BPanopolis帕诺波利斯
	f.帕诺波利斯Panopolis.SetParent(f)
	
	f.祖赫拉Zuhrah = BZuhrah祖赫拉
	f.祖赫拉Zuhrah.SetParent(f)
	
}
