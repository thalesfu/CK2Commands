package kambampet

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KambampetCounty interface {
	feud.County
	BBhadrachalam跋陀罗遮蓝() feud.Barony
	BKambampet剑婆摩波底() feud.Barony
	BMadupalli马杜帕利() feud.Barony
	BNelakondapalli泥罗军荼波梨() feud.Barony
	BPenuballi佩努巴利() feud.Barony
	BShahpura沙普拉() feud.Barony
	BSuryapet苏利耶佩特() feud.Barony
}

type 剑婆摩波底KambampetCounty struct {
	feud.BaseCounty
	跋陀罗遮蓝Bhadrachalam    feud.Barony
	剑婆摩波底Kambampet       feud.Barony
	马杜帕利Madupalli        feud.Barony
	泥罗军荼波梨Nelakondapalli feud.Barony
	佩努巴利Penuballi        feud.Barony
	沙普拉Shahpura          feud.Barony
	苏利耶佩特Suryapet        feud.Barony
}

func (c *剑婆摩波底KambampetCounty) BBhadrachalam跋陀罗遮蓝() feud.Barony {
	return c.跋陀罗遮蓝Bhadrachalam
}

func (c *剑婆摩波底KambampetCounty) BKambampet剑婆摩波底() feud.Barony {
	return c.剑婆摩波底Kambampet
}

func (c *剑婆摩波底KambampetCounty) BMadupalli马杜帕利() feud.Barony {
	return c.马杜帕利Madupalli
}

func (c *剑婆摩波底KambampetCounty) BNelakondapalli泥罗军荼波梨() feud.Barony {
	return c.泥罗军荼波梨Nelakondapalli
}

func (c *剑婆摩波底KambampetCounty) BPenuballi佩努巴利() feud.Barony {
	return c.佩努巴利Penuballi
}

func (c *剑婆摩波底KambampetCounty) BShahpura沙普拉() feud.Barony {
	return c.沙普拉Shahpura
}

func (c *剑婆摩波底KambampetCounty) BSuryapet苏利耶佩特() feud.Barony {
	return c.苏利耶佩特Suryapet
}

var CKambampet剑婆摩波底 KambampetCounty = &剑婆摩波底KambampetCounty{}

func init() {
	f := CKambampet剑婆摩波底.(*剑婆摩波底KambampetCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1208",
		Title:     "kambampet",
		TitleName: "剑婆摩波底",
		TitleCode: "c_kambampet",
		Baronies:  map[string]feud.Barony{},
	}

	f.跋陀罗遮蓝Bhadrachalam = BBhadrachalam跋陀罗遮蓝
	f.跋陀罗遮蓝Bhadrachalam.SetParent(f)

	f.剑婆摩波底Kambampet = BKambampet剑婆摩波底
	f.剑婆摩波底Kambampet.SetParent(f)

	f.马杜帕利Madupalli = BMadupalli马杜帕利
	f.马杜帕利Madupalli.SetParent(f)

	f.泥罗军荼波梨Nelakondapalli = BNelakondapalli泥罗军荼波梨
	f.泥罗军荼波梨Nelakondapalli.SetParent(f)

	f.佩努巴利Penuballi = BPenuballi佩努巴利
	f.佩努巴利Penuballi.SetParent(f)

	f.沙普拉Shahpura = BShahpura沙普拉
	f.沙普拉Shahpura.SetParent(f)

	f.苏利耶佩特Suryapet = BSuryapet苏利耶佩特
	f.苏利耶佩特Suryapet.SetParent(f)

}
