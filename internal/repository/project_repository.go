package repository

import (
	"context"

	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IProjectRepository interface {
	GetProjects(userId string) ([]model.Project, error)
	CreateProject(token model.Project) error
}

type ProjectRepository struct {
	ctx    *context.Context
	dbPool *pgxpool.Pool
}

func NewProjectRepository(ctx context.Context, dbPool *pgxpool.Pool) IProjectRepository {
	return &ProjectRepository{ctx: &ctx, dbPool: dbPool}
}

// CreateToken implements ITokenRepository.
func (repository *ProjectRepository) CreateProject(project model.Project) error {
	_, err := repository.dbPool.Exec(*repository.ctx,
		"INSERT INTO project (id,user_id,name) VALUES ($1,$2,$3)",
		project.Id, project.UserId, project.Name,
	)
	return err
}

// GetTokenByUserId implements ITokenRepository.
func (repository *ProjectRepository) GetProjects(userId string) ([]model.Project, error) {
	var projects []model.Project

	row, err := repository.dbPool.Query(*repository.ctx,
		"SELECT * FROM projects WHERE user_id = $1", userId,
	)

	var project model.Project
	for row.Next() {
		err = row.Scan(&project.Id, &project.UserId, &project.Name)
		if err != nil {
			break
		}
		projects = append(projects, project)
	}

	return projects, err
}
