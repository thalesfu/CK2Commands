package sagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SagarCounty interface {
	feud.County
	BBagavadi巴加瓦迪() feud.Barony
	BBagpura薄伽补罗() feud.Barony
	BBandaguda槃陀瞿荼() feud.Barony
	BPattadakal帕塔达卡尔() feud.Barony
	BSagar娑伽罗() feud.Barony
	BSorapur苏拉补罗() feud.Barony
	BTalikota达利戈达() feud.Barony
}

type 娑伽罗SagarCounty struct {
	feud.BaseCounty
	巴加瓦迪Bagavadi    feud.Barony
	薄伽补罗Bagpura     feud.Barony
	槃陀瞿荼Bandaguda   feud.Barony
	帕塔达卡尔Pattadakal feud.Barony
	娑伽罗Sagar        feud.Barony
	苏拉补罗Sorapur     feud.Barony
	达利戈达Talikota    feud.Barony
}

func (c *娑伽罗SagarCounty) BBagavadi巴加瓦迪() feud.Barony {
	return c.巴加瓦迪Bagavadi
}

func (c *娑伽罗SagarCounty) BBagpura薄伽补罗() feud.Barony {
	return c.薄伽补罗Bagpura
}

func (c *娑伽罗SagarCounty) BBandaguda槃陀瞿荼() feud.Barony {
	return c.槃陀瞿荼Bandaguda
}

func (c *娑伽罗SagarCounty) BPattadakal帕塔达卡尔() feud.Barony {
	return c.帕塔达卡尔Pattadakal
}

func (c *娑伽罗SagarCounty) BSagar娑伽罗() feud.Barony {
	return c.娑伽罗Sagar
}

func (c *娑伽罗SagarCounty) BSorapur苏拉补罗() feud.Barony {
	return c.苏拉补罗Sorapur
}

func (c *娑伽罗SagarCounty) BTalikota达利戈达() feud.Barony {
	return c.达利戈达Talikota
}

var CSagar娑伽罗 SagarCounty = &娑伽罗SagarCounty{}

func init() {
	f := CSagar娑伽罗.(*娑伽罗SagarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1265",
		Title:     "sagar",
		TitleName: "娑伽罗",
		TitleCode: "c_sagar",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴加瓦迪Bagavadi = BBagavadi巴加瓦迪
	f.巴加瓦迪Bagavadi.SetParent(f)

	f.薄伽补罗Bagpura = BBagpura薄伽补罗
	f.薄伽补罗Bagpura.SetParent(f)

	f.槃陀瞿荼Bandaguda = BBandaguda槃陀瞿荼
	f.槃陀瞿荼Bandaguda.SetParent(f)

	f.帕塔达卡尔Pattadakal = BPattadakal帕塔达卡尔
	f.帕塔达卡尔Pattadakal.SetParent(f)

	f.娑伽罗Sagar = BSagar娑伽罗
	f.娑伽罗Sagar.SetParent(f)

	f.苏拉补罗Sorapur = BSorapur苏拉补罗
	f.苏拉补罗Sorapur.SetParent(f)

	f.达利戈达Talikota = BTalikota达利戈达
	f.达利戈达Talikota.SetParent(f)

}
