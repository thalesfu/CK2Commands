package armail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArmailCounty interface {
    feud.County
    BArmail亚梅尔() 	feud.Barony
    BBasti_basira巴斯蒂巴茜拉() 	feud.Barony
    BDaryun达龙() 	feud.Barony
    BHingula辛古拉() 	feud.Barony
    BKambali坎巴里() 	feud.Barony
    BPejoke普咎柯() 	feud.Barony
    BRihana_sahu蕾哈纳沙笏() 	feud.Barony
    BYusli育里() 	feud.Barony
}

type 阿尔梅尔ArmailCounty struct {
	feud.BaseCounty
	亚梅尔Armail 	feud.Barony
	巴斯蒂巴茜拉Basti_basira 	feud.Barony
	达龙Daryun 	feud.Barony
	辛古拉Hingula 	feud.Barony
	坎巴里Kambali 	feud.Barony
	普咎柯Pejoke 	feud.Barony
	蕾哈纳沙笏Rihana_sahu 	feud.Barony
	育里Yusli 	feud.Barony
}

func (c *阿尔梅尔ArmailCounty) BArmail亚梅尔() feud.Barony {
	return c.亚梅尔Armail
}
    
func (c *阿尔梅尔ArmailCounty) BBasti_basira巴斯蒂巴茜拉() feud.Barony {
	return c.巴斯蒂巴茜拉Basti_basira
}
    
func (c *阿尔梅尔ArmailCounty) BDaryun达龙() feud.Barony {
	return c.达龙Daryun
}
    
func (c *阿尔梅尔ArmailCounty) BHingula辛古拉() feud.Barony {
	return c.辛古拉Hingula
}
    
func (c *阿尔梅尔ArmailCounty) BKambali坎巴里() feud.Barony {
	return c.坎巴里Kambali
}
    
func (c *阿尔梅尔ArmailCounty) BPejoke普咎柯() feud.Barony {
	return c.普咎柯Pejoke
}
    
func (c *阿尔梅尔ArmailCounty) BRihana_sahu蕾哈纳沙笏() feud.Barony {
	return c.蕾哈纳沙笏Rihana_sahu
}
    
func (c *阿尔梅尔ArmailCounty) BYusli育里() feud.Barony {
	return c.育里Yusli
}
    
var CArmail阿尔梅尔 ArmailCounty = &阿尔梅尔ArmailCounty{}

func init() {
	f := CArmail阿尔梅尔.(*阿尔梅尔ArmailCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1370",
		Title:     "armail",
		TitleName: "阿尔梅尔",
		TitleCode: "c_armail",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚梅尔Armail = BArmail亚梅尔
	f.亚梅尔Armail.SetParent(f)
	
	f.巴斯蒂巴茜拉Basti_basira = BBasti_basira巴斯蒂巴茜拉
	f.巴斯蒂巴茜拉Basti_basira.SetParent(f)
	
	f.达龙Daryun = BDaryun达龙
	f.达龙Daryun.SetParent(f)
	
	f.辛古拉Hingula = BHingula辛古拉
	f.辛古拉Hingula.SetParent(f)
	
	f.坎巴里Kambali = BKambali坎巴里
	f.坎巴里Kambali.SetParent(f)
	
	f.普咎柯Pejoke = BPejoke普咎柯
	f.普咎柯Pejoke.SetParent(f)
	
	f.蕾哈纳沙笏Rihana_sahu = BRihana_sahu蕾哈纳沙笏
	f.蕾哈纳沙笏Rihana_sahu.SetParent(f)
	
	f.育里Yusli = BYusli育里
	f.育里Yusli.SetParent(f)
	
}
