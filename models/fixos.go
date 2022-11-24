package models

type Fixos struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Contas      float64 `json:"contas"`
	Assinaturas float64 `json:"assinaturas"`
	Seguros     float64 `json:"seguros"`
	Mesadas     float64 `json:"mesadas"`
	Impostos    float64 `json:"impostos"`
	Outros      float64 `json:"outros"`
	Mes         int     `json:"mes"`
	Ano         int     `json:"ano"`
}
