package dto

type ReqPatient struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Indication string `json:"indication"`
}
