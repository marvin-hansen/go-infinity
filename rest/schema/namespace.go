package schema

import "fmt"

type Namespace struct {
	Name       string   `json:"name"`
	Classes    []Schema `json:"classes"`
	RawMessage []byte
}

func (s *Namespace) String() string {
	if s == nil {
		return ""
	} else {
		return fmt.Sprintf("\n Name: %v \n Classes: %v \n  ",
			s.Name,
			s.Classes,
		)
	}
}
