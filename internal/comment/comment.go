package comment

import (
	"github.com/jinzhu/gorm"
)

// Service - service structfor our comment service
type Service struct {
	DB *gorm.DB
}

// Comment - defiines our comments structure
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

// CommentService - the interface for our comment service
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComents() ([]Comment, error)
}

// NewService - returns a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetComment - returns comment from DB by its ID
func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// GetCommentBySlug - returns comments by slug
func (s *Service) GetCommentBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments).Where("slug=?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

// PostComment - save comment to DB and return it's entity if successful
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}

	if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

// DeleteComment - deletes a comment from the database by ID
func (s *Service) DeleteComment(ID uint) error {
	var comment Comment
	if result := s.DB.Delete(&comment, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllComents - retrievers all comments from the database
func (s *Service) GetAllComents() ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}
