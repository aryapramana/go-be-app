package mapping

type SourceNews struct {
	ID   string `json:"source_id"`
	Name string `json:"source_name"`
}

func (sourcenews *SourceNews) SetId(ID string) {
	sourcenews.ID = ID
}

func (sourcenews SourceNews) GetId() string {
	return sourcenews.ID
}

func (sourcenews *SourceNews) SetName(Name string) {
	sourcenews.Name = Name
}

func (sourcenews SourceNews) GetName() string {
	return sourcenews.Name
}
