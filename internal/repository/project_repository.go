package repository

import (
	"context"

	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IProjectRepository interface {
	GetProjectReportByUserId(userId string) ([]dto.ListProjectResponse, error)
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
		"INSERT INTO projects (project_id,user_id,name,created_at) VALUES ($1,$2,$3,$4)",
		project.ProjectId, project.UserId, project.Name, project.CreatedAt,
	)
	return err
}

// GetTokenByUserId implements ITokenRepository.
func (repository *ProjectRepository) GetProjectReportByUserId(userId string) ([]dto.ListProjectResponse, error) {
	var projects []dto.ListProjectResponse

	row, err := repository.dbPool.Query(*repository.ctx,
		`
		SELECT
		p.project_id,
		COALESCE(p.name, '') AS project_name,
		COALESCE(SUM(e.amount), 0) AS total_amount,
		p.created_at AS created_at,
		COALESCE(COUNT(e.expense_id), 0) AS total_expenses
	FROM
		projects p
	LEFT JOIN
		expenses e ON p.project_id = e.project_id
	WHERE 
		p.user_id = $1
	GROUP BY
    p.project_id;
		`, userId,
	)

	var project dto.ListProjectResponse
	for row.Next() {
		err = row.Scan(&project.ProjectId, &project.ProjectName, &project.TotalAmount, &project.CreatedAt, &project.TotalExpenses)
		if err != nil {
			break
		}
		projects = append(projects, project)
	}

	return projects, err
}
