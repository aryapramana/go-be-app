package model

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (source *Source) SetID(ID string) {
	source.ID = ID
}

func (source Source) GetID() string {
	return source.ID
}

func (source *Source) SetName(Name string) {
	source.Name = Name
}

func (source Source) GetName() string {
	return source.Name
}
