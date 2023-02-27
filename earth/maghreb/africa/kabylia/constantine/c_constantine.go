package constantine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ConstantineCounty interface {
    feud.County
    BConstantine君士坦丁() 	feud.Barony
    BElkhroub赫鲁卜() 	feud.Barony
    BGuelma盖尔马() 	feud.Barony
    BGuerrah古尔拉哈() 	feud.Barony
    BMechtakebira梅彻塔科比拉() 	feud.Barony
    BOumelbouaghi乌姆布阿基() 	feud.Barony
    BTijis提吉斯() 	feud.Barony
}

type 君士坦丁ConstantineCounty struct {
	feud.BaseCounty
	君士坦丁Constantine 	feud.Barony
	赫鲁卜Elkhroub 	feud.Barony
	盖尔马Guelma 	feud.Barony
	古尔拉哈Guerrah 	feud.Barony
	梅彻塔科比拉Mechtakebira 	feud.Barony
	乌姆布阿基Oumelbouaghi 	feud.Barony
	提吉斯Tijis 	feud.Barony
}

func (c *君士坦丁ConstantineCounty) BConstantine君士坦丁() feud.Barony {
	return c.君士坦丁Constantine
}
    
func (c *君士坦丁ConstantineCounty) BElkhroub赫鲁卜() feud.Barony {
	return c.赫鲁卜Elkhroub
}
    
func (c *君士坦丁ConstantineCounty) BGuelma盖尔马() feud.Barony {
	return c.盖尔马Guelma
}
    
func (c *君士坦丁ConstantineCounty) BGuerrah古尔拉哈() feud.Barony {
	return c.古尔拉哈Guerrah
}
    
func (c *君士坦丁ConstantineCounty) BMechtakebira梅彻塔科比拉() feud.Barony {
	return c.梅彻塔科比拉Mechtakebira
}
    
func (c *君士坦丁ConstantineCounty) BOumelbouaghi乌姆布阿基() feud.Barony {
	return c.乌姆布阿基Oumelbouaghi
}
    
func (c *君士坦丁ConstantineCounty) BTijis提吉斯() feud.Barony {
	return c.提吉斯Tijis
}
    
var CConstantine君士坦丁 ConstantineCounty = &君士坦丁ConstantineCounty{}

func init() {
	f := CConstantine君士坦丁.(*君士坦丁ConstantineCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "821",
		Title:     "constantine",
		TitleName: "君士坦丁",
		TitleCode: "c_constantine",
		Baronies:  map[string]feud.Barony{},
	}

	f.君士坦丁Constantine = BConstantine君士坦丁
	f.君士坦丁Constantine.SetParent(f)
	
	f.赫鲁卜Elkhroub = BElkhroub赫鲁卜
	f.赫鲁卜Elkhroub.SetParent(f)
	
	f.盖尔马Guelma = BGuelma盖尔马
	f.盖尔马Guelma.SetParent(f)
	
	f.古尔拉哈Guerrah = BGuerrah古尔拉哈
	f.古尔拉哈Guerrah.SetParent(f)
	
	f.梅彻塔科比拉Mechtakebira = BMechtakebira梅彻塔科比拉
	f.梅彻塔科比拉Mechtakebira.SetParent(f)
	
	f.乌姆布阿基Oumelbouaghi = BOumelbouaghi乌姆布阿基
	f.乌姆布阿基Oumelbouaghi.SetParent(f)
	
	f.提吉斯Tijis = BTijis提吉斯
	f.提吉斯Tijis.SetParent(f)
	
}
