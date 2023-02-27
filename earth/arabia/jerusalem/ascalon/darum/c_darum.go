package darum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DarumCounty interface {
    feud.County
    BAbasan阿巴桑() 	feud.Barony
    BDarum达鲁姆() 	feud.Barony
    BGaza加沙() 	feud.Barony
    BGerar基拉耳() 	feud.Barony
    BJarara贾雷拉() 	feud.Barony
    BRadwan拉德万() 	feud.Barony
    BRafah拉法() 	feud.Barony
    BSalqah塞奇赫() 	feud.Barony
}

type 达鲁姆DarumCounty struct {
	feud.BaseCounty
	阿巴桑Abasan 	feud.Barony
	达鲁姆Darum 	feud.Barony
	加沙Gaza 	feud.Barony
	基拉耳Gerar 	feud.Barony
	贾雷拉Jarara 	feud.Barony
	拉德万Radwan 	feud.Barony
	拉法Rafah 	feud.Barony
	塞奇赫Salqah 	feud.Barony
}

func (c *达鲁姆DarumCounty) BAbasan阿巴桑() feud.Barony {
	return c.阿巴桑Abasan
}
    
func (c *达鲁姆DarumCounty) BDarum达鲁姆() feud.Barony {
	return c.达鲁姆Darum
}
    
func (c *达鲁姆DarumCounty) BGaza加沙() feud.Barony {
	return c.加沙Gaza
}
    
func (c *达鲁姆DarumCounty) BGerar基拉耳() feud.Barony {
	return c.基拉耳Gerar
}
    
func (c *达鲁姆DarumCounty) BJarara贾雷拉() feud.Barony {
	return c.贾雷拉Jarara
}
    
func (c *达鲁姆DarumCounty) BRadwan拉德万() feud.Barony {
	return c.拉德万Radwan
}
    
func (c *达鲁姆DarumCounty) BRafah拉法() feud.Barony {
	return c.拉法Rafah
}
    
func (c *达鲁姆DarumCounty) BSalqah塞奇赫() feud.Barony {
	return c.塞奇赫Salqah
}
    
var CDarum达鲁姆 DarumCounty = &达鲁姆DarumCounty{}

func init() {
	f := CDarum达鲁姆.(*达鲁姆DarumCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "781",
		Title:     "darum",
		TitleName: "达鲁姆",
		TitleCode: "c_darum",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴桑Abasan = BAbasan阿巴桑
	f.阿巴桑Abasan.SetParent(f)
	
	f.达鲁姆Darum = BDarum达鲁姆
	f.达鲁姆Darum.SetParent(f)
	
	f.加沙Gaza = BGaza加沙
	f.加沙Gaza.SetParent(f)
	
	f.基拉耳Gerar = BGerar基拉耳
	f.基拉耳Gerar.SetParent(f)
	
	f.贾雷拉Jarara = BJarara贾雷拉
	f.贾雷拉Jarara.SetParent(f)
	
	f.拉德万Radwan = BRadwan拉德万
	f.拉德万Radwan.SetParent(f)
	
	f.拉法Rafah = BRafah拉法
	f.拉法Rafah.SetParent(f)
	
	f.塞奇赫Salqah = BSalqah塞奇赫
	f.塞奇赫Salqah.SetParent(f)
	
}
