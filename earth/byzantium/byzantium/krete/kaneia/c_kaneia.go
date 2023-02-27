package kaneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KaneiaCounty interface {
    feud.County
    BAkrotin阿克洛汀() 	feud.Barony
    BArkadi阿尔卡季() 	feud.Barony
    BKandia坎迪亚() 	feud.Barony
    BKastellikissamos卡斯泰利_基萨莫斯() 	feud.Barony
    BMatala马塔拉() 	feud.Barony
    BNikiforosfokas尼基佛洛佛卡斯() 	feud.Barony
    BPaleohora帕莱奥霍拉() 	feud.Barony
    BRethymno雷西姆诺() 	feud.Barony
}

type 卡尼雅KaneiaCounty struct {
	feud.BaseCounty
	阿克洛汀Akrotin 	feud.Barony
	阿尔卡季Arkadi 	feud.Barony
	坎迪亚Kandia 	feud.Barony
	卡斯泰利_基萨莫斯Kastellikissamos 	feud.Barony
	马塔拉Matala 	feud.Barony
	尼基佛洛佛卡斯Nikiforosfokas 	feud.Barony
	帕莱奥霍拉Paleohora 	feud.Barony
	雷西姆诺Rethymno 	feud.Barony
}

func (c *卡尼雅KaneiaCounty) BAkrotin阿克洛汀() feud.Barony {
	return c.阿克洛汀Akrotin
}
    
func (c *卡尼雅KaneiaCounty) BArkadi阿尔卡季() feud.Barony {
	return c.阿尔卡季Arkadi
}
    
func (c *卡尼雅KaneiaCounty) BKandia坎迪亚() feud.Barony {
	return c.坎迪亚Kandia
}
    
func (c *卡尼雅KaneiaCounty) BKastellikissamos卡斯泰利_基萨莫斯() feud.Barony {
	return c.卡斯泰利_基萨莫斯Kastellikissamos
}
    
func (c *卡尼雅KaneiaCounty) BMatala马塔拉() feud.Barony {
	return c.马塔拉Matala
}
    
func (c *卡尼雅KaneiaCounty) BNikiforosfokas尼基佛洛佛卡斯() feud.Barony {
	return c.尼基佛洛佛卡斯Nikiforosfokas
}
    
func (c *卡尼雅KaneiaCounty) BPaleohora帕莱奥霍拉() feud.Barony {
	return c.帕莱奥霍拉Paleohora
}
    
func (c *卡尼雅KaneiaCounty) BRethymno雷西姆诺() feud.Barony {
	return c.雷西姆诺Rethymno
}
    
var CKaneia卡尼雅 KaneiaCounty = &卡尼雅KaneiaCounty{}

func init() {
	f := CKaneia卡尼雅.(*卡尼雅KaneiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "479",
		Title:     "kaneia",
		TitleName: "卡尼雅",
		TitleCode: "c_kaneia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克洛汀Akrotin = BAkrotin阿克洛汀
	f.阿克洛汀Akrotin.SetParent(f)
	
	f.阿尔卡季Arkadi = BArkadi阿尔卡季
	f.阿尔卡季Arkadi.SetParent(f)
	
	f.坎迪亚Kandia = BKandia坎迪亚
	f.坎迪亚Kandia.SetParent(f)
	
	f.卡斯泰利_基萨莫斯Kastellikissamos = BKastellikissamos卡斯泰利_基萨莫斯
	f.卡斯泰利_基萨莫斯Kastellikissamos.SetParent(f)
	
	f.马塔拉Matala = BMatala马塔拉
	f.马塔拉Matala.SetParent(f)
	
	f.尼基佛洛佛卡斯Nikiforosfokas = BNikiforosfokas尼基佛洛佛卡斯
	f.尼基佛洛佛卡斯Nikiforosfokas.SetParent(f)
	
	f.帕莱奥霍拉Paleohora = BPaleohora帕莱奥霍拉
	f.帕莱奥霍拉Paleohora.SetParent(f)
	
	f.雷西姆诺Rethymno = BRethymno雷西姆诺
	f.雷西姆诺Rethymno.SetParent(f)
	
}
