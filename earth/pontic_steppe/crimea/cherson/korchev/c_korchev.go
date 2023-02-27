package korchev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KorchevCounty interface {
    feud.County
    BBaherove巴格罗韦() 	feud.Barony
    BBosphoros博斯普鲁斯() 	feud.Barony
    BCherco切尔格() 	feud.Barony
    BChystopillia基斯托比利亚() 	feud.Barony
    BNymphaion尼姆法延() 	feud.Barony
    BPanticapea潘提卡派翁() 	feud.Barony
    BVosporo沃斯波罗() 	feud.Barony
    BZavitne扎维特尼() 	feud.Barony
}

type 刻赤KorchevCounty struct {
	feud.BaseCounty
	巴格罗韦Baherove 	feud.Barony
	博斯普鲁斯Bosphoros 	feud.Barony
	切尔格Cherco 	feud.Barony
	基斯托比利亚Chystopillia 	feud.Barony
	尼姆法延Nymphaion 	feud.Barony
	潘提卡派翁Panticapea 	feud.Barony
	沃斯波罗Vosporo 	feud.Barony
	扎维特尼Zavitne 	feud.Barony
}

func (c *刻赤KorchevCounty) BBaherove巴格罗韦() feud.Barony {
	return c.巴格罗韦Baherove
}
    
func (c *刻赤KorchevCounty) BBosphoros博斯普鲁斯() feud.Barony {
	return c.博斯普鲁斯Bosphoros
}
    
func (c *刻赤KorchevCounty) BCherco切尔格() feud.Barony {
	return c.切尔格Cherco
}
    
func (c *刻赤KorchevCounty) BChystopillia基斯托比利亚() feud.Barony {
	return c.基斯托比利亚Chystopillia
}
    
func (c *刻赤KorchevCounty) BNymphaion尼姆法延() feud.Barony {
	return c.尼姆法延Nymphaion
}
    
func (c *刻赤KorchevCounty) BPanticapea潘提卡派翁() feud.Barony {
	return c.潘提卡派翁Panticapea
}
    
func (c *刻赤KorchevCounty) BVosporo沃斯波罗() feud.Barony {
	return c.沃斯波罗Vosporo
}
    
func (c *刻赤KorchevCounty) BZavitne扎维特尼() feud.Barony {
	return c.扎维特尼Zavitne
}
    
var CKorchev刻赤 KorchevCounty = &刻赤KorchevCounty{}

func init() {
	f := CKorchev刻赤.(*刻赤KorchevCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "562",
		Title:     "korchev",
		TitleName: "刻赤",
		TitleCode: "c_korchev",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴格罗韦Baherove = BBaherove巴格罗韦
	f.巴格罗韦Baherove.SetParent(f)
	
	f.博斯普鲁斯Bosphoros = BBosphoros博斯普鲁斯
	f.博斯普鲁斯Bosphoros.SetParent(f)
	
	f.切尔格Cherco = BCherco切尔格
	f.切尔格Cherco.SetParent(f)
	
	f.基斯托比利亚Chystopillia = BChystopillia基斯托比利亚
	f.基斯托比利亚Chystopillia.SetParent(f)
	
	f.尼姆法延Nymphaion = BNymphaion尼姆法延
	f.尼姆法延Nymphaion.SetParent(f)
	
	f.潘提卡派翁Panticapea = BPanticapea潘提卡派翁
	f.潘提卡派翁Panticapea.SetParent(f)
	
	f.沃斯波罗Vosporo = BVosporo沃斯波罗
	f.沃斯波罗Vosporo.SetParent(f)
	
	f.扎维特尼Zavitne = BZavitne扎维特尼
	f.扎维特尼Zavitne.SetParent(f)
	
}
