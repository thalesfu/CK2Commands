package lower_dniepr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Lower_dnieprCounty interface {
    feud.County
    BBilozerka别洛焦尔卡() 	feud.Barony
    BKairy卡伊里() 	feud.Barony
    BKuturogly库图尔_奥格雷() 	feud.Barony
    BKyzylyar克孜勒亚尔() 	feud.Barony
    BOleshye奥列什基() 	feud.Barony
    BPryazovske普里阿佐夫西克() 	feud.Barony
    BShagingirei沙金_吉雷() 	feud.Barony
    BTokmak托克马克() 	feud.Barony
}

type 奥列什基Lower_dnieprCounty struct {
	feud.BaseCounty
	别洛焦尔卡Bilozerka 	feud.Barony
	卡伊里Kairy 	feud.Barony
	库图尔_奥格雷Kuturogly 	feud.Barony
	克孜勒亚尔Kyzylyar 	feud.Barony
	奥列什基Oleshye 	feud.Barony
	普里阿佐夫西克Pryazovske 	feud.Barony
	沙金_吉雷Shagingirei 	feud.Barony
	托克马克Tokmak 	feud.Barony
}

func (c *奥列什基Lower_dnieprCounty) BBilozerka别洛焦尔卡() feud.Barony {
	return c.别洛焦尔卡Bilozerka
}
    
func (c *奥列什基Lower_dnieprCounty) BKairy卡伊里() feud.Barony {
	return c.卡伊里Kairy
}
    
func (c *奥列什基Lower_dnieprCounty) BKuturogly库图尔_奥格雷() feud.Barony {
	return c.库图尔_奥格雷Kuturogly
}
    
func (c *奥列什基Lower_dnieprCounty) BKyzylyar克孜勒亚尔() feud.Barony {
	return c.克孜勒亚尔Kyzylyar
}
    
func (c *奥列什基Lower_dnieprCounty) BOleshye奥列什基() feud.Barony {
	return c.奥列什基Oleshye
}
    
func (c *奥列什基Lower_dnieprCounty) BPryazovske普里阿佐夫西克() feud.Barony {
	return c.普里阿佐夫西克Pryazovske
}
    
func (c *奥列什基Lower_dnieprCounty) BShagingirei沙金_吉雷() feud.Barony {
	return c.沙金_吉雷Shagingirei
}
    
func (c *奥列什基Lower_dnieprCounty) BTokmak托克马克() feud.Barony {
	return c.托克马克Tokmak
}
    
var CLower_dniepr奥列什基 Lower_dnieprCounty = &奥列什基Lower_dnieprCounty{}

func init() {
	f := CLower_dniepr奥列什基.(*奥列什基Lower_dnieprCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "558",
		Title:     "lower_dniepr",
		TitleName: "奥列什基",
		TitleCode: "c_lower_dniepr",
		Baronies:  map[string]feud.Barony{},
	}

	f.别洛焦尔卡Bilozerka = BBilozerka别洛焦尔卡
	f.别洛焦尔卡Bilozerka.SetParent(f)
	
	f.卡伊里Kairy = BKairy卡伊里
	f.卡伊里Kairy.SetParent(f)
	
	f.库图尔_奥格雷Kuturogly = BKuturogly库图尔_奥格雷
	f.库图尔_奥格雷Kuturogly.SetParent(f)
	
	f.克孜勒亚尔Kyzylyar = BKyzylyar克孜勒亚尔
	f.克孜勒亚尔Kyzylyar.SetParent(f)
	
	f.奥列什基Oleshye = BOleshye奥列什基
	f.奥列什基Oleshye.SetParent(f)
	
	f.普里阿佐夫西克Pryazovske = BPryazovske普里阿佐夫西克
	f.普里阿佐夫西克Pryazovske.SetParent(f)
	
	f.沙金_吉雷Shagingirei = BShagingirei沙金_吉雷
	f.沙金_吉雷Shagingirei.SetParent(f)
	
	f.托克马克Tokmak = BTokmak托克马克
	f.托克马克Tokmak.SetParent(f)
	
}
