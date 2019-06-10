package repository

import (
  "CPEProject/src/models"
)

//ProfileRepository interface
type FacultyRepository interface {
  Save(*models.Faculty) error
  Update(string, *models.Faculty) error
  Delete(string) error
  FindByID(string) (*models.Faculty, error)
  SaveTest(*models.Faculty) error
  //FindAll() (models.Profiles, error)
}
type MajorRepository interface {
  Save(*models.Major) error
  //Update(string, *models.Faculty) error
  //Delete(string) error
  //FindByID(string) (*models.Faculty, error)
  //FindAll() (models.Profiles, error)
}

type SubjectRepository interface {
  Save(*models.Subject) error
  //Update(string, *models.Faculty) error
  //Delete(string) error
  //FindByID(string) (*models.Faculty, error)
  //FindAll() (models.Profiles, error)
}