package models

type Student struct {
	// Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Student_id   string `json:"student_id,omitempty"`
	Student_name string `json:"student_name,omitempty"`
	Class        string `json:"class,omitempty"`
	Marks        int    `json:"marks,omitempty"`
}
