package validation

import "github.com/google/uuid"

func UUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return err
	}
	return nil
}
