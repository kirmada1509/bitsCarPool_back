package models

type Notificaton struct{
	From string
	To string
	Trip_id string
	Type string //"request" | "accept" | "decline"
}