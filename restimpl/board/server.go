package board

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/shinsuke-nara/echo-server/db"
	"github.com/shinsuke-nara/echo-server/restapi/operations/board"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) PostMessage(params board.PostMessageParams) middleware.Responder {
	result := s.DB.Create(&db.Message{Body: *params.Massage.Message})
	if result.Error != nil {
		return board.NewPostMessageInternalServerError().WithPayload(result.Error.Error())
	}
	return &board.PostMessageOK{}
}

func (s *Server) ListMessage(params board.ListMessageParams) middleware.Responder {
	messages := []db.Message{}
	result := s.DB.Find(&messages)
	if result.Error != nil {
		return board.NewListMessageInternalServerError().WithPayload(result.Error.Error())
	}
	items := make([]*board.ListMessageOKBodyItems0, 0, len(messages))
	for _, m := range messages {
		id := int64(m.ID)
		items = append(items, &board.ListMessageOKBodyItems0{
			ID:      &id,
			Message: &m.Body,
		})
	}
	return board.NewListMessageOK().WithPayload(items)
}
