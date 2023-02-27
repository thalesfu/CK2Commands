package kandalax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KandalaxCounty interface {
    feud.County
    BKantalahti坎塔拉赫蒂() 	feud.Barony
    BKolsky科尔斯基() 	feud.Barony
    BLekastrom列卡斯特罗姆() 	feud.Barony
    BPonoy波诺伊() 	feud.Barony
    BPyaozero皮亚奥泽罗() 	feud.Barony
    BSarapo萨拉帕() 	feud.Barony
    BUmba温巴() 	feud.Barony
    BVarzuga瓦尔祖加() 	feud.Barony
}

type 坎达拉克斯KandalaxCounty struct {
	feud.BaseCounty
	坎塔拉赫蒂Kantalahti 	feud.Barony
	科尔斯基Kolsky 	feud.Barony
	列卡斯特罗姆Lekastrom 	feud.Barony
	波诺伊Ponoy 	feud.Barony
	皮亚奥泽罗Pyaozero 	feud.Barony
	萨拉帕Sarapo 	feud.Barony
	温巴Umba 	feud.Barony
	瓦尔祖加Varzuga 	feud.Barony
}

func (c *坎达拉克斯KandalaxCounty) BKantalahti坎塔拉赫蒂() feud.Barony {
	return c.坎塔拉赫蒂Kantalahti
}
    
func (c *坎达拉克斯KandalaxCounty) BKolsky科尔斯基() feud.Barony {
	return c.科尔斯基Kolsky
}
    
func (c *坎达拉克斯KandalaxCounty) BLekastrom列卡斯特罗姆() feud.Barony {
	return c.列卡斯特罗姆Lekastrom
}
    
func (c *坎达拉克斯KandalaxCounty) BPonoy波诺伊() feud.Barony {
	return c.波诺伊Ponoy
}
    
func (c *坎达拉克斯KandalaxCounty) BPyaozero皮亚奥泽罗() feud.Barony {
	return c.皮亚奥泽罗Pyaozero
}
    
func (c *坎达拉克斯KandalaxCounty) BSarapo萨拉帕() feud.Barony {
	return c.萨拉帕Sarapo
}
    
func (c *坎达拉克斯KandalaxCounty) BUmba温巴() feud.Barony {
	return c.温巴Umba
}
    
func (c *坎达拉克斯KandalaxCounty) BVarzuga瓦尔祖加() feud.Barony {
	return c.瓦尔祖加Varzuga
}
    
var CKandalax坎达拉克斯 KandalaxCounty = &坎达拉克斯KandalaxCounty{}

func init() {
	f := CKandalax坎达拉克斯.(*坎达拉克斯KandalaxCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "850",
		Title:     "kandalax",
		TitleName: "坎达拉克斯",
		TitleCode: "c_kandalax",
		Baronies:  map[string]feud.Barony{},
	}

	f.坎塔拉赫蒂Kantalahti = BKantalahti坎塔拉赫蒂
	f.坎塔拉赫蒂Kantalahti.SetParent(f)
	
	f.科尔斯基Kolsky = BKolsky科尔斯基
	f.科尔斯基Kolsky.SetParent(f)
	
	f.列卡斯特罗姆Lekastrom = BLekastrom列卡斯特罗姆
	f.列卡斯特罗姆Lekastrom.SetParent(f)
	
	f.波诺伊Ponoy = BPonoy波诺伊
	f.波诺伊Ponoy.SetParent(f)
	
	f.皮亚奥泽罗Pyaozero = BPyaozero皮亚奥泽罗
	f.皮亚奥泽罗Pyaozero.SetParent(f)
	
	f.萨拉帕Sarapo = BSarapo萨拉帕
	f.萨拉帕Sarapo.SetParent(f)
	
	f.温巴Umba = BUmba温巴
	f.温巴Umba.SetParent(f)
	
	f.瓦尔祖加Varzuga = BVarzuga瓦尔祖加
	f.瓦尔祖加Varzuga.SetParent(f)
	
}
