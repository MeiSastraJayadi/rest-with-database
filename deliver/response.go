package deliver

import (
	"io"

	"github.com/MeiSastraJayadi/rest-with-datatabase/usecase"
)

type Reponse struct {
	message string `json:"message"`
}

func (c *Reponse) Length() int {
	return len(c.message)
}

func (c *Reponse) WriteMessage(message string) {
	c.message = message
}

func (c *Reponse) JSONResponse(w io.Writer) error {
	err := usecase.ToJSON(w, c)
	if err != nil {
		return err
	}
	return nil
}
