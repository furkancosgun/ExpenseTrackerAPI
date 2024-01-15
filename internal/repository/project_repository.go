package repository

import (
	"context"

	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IProjectRepository interface {
	GetProjectReportByUserId(userId string) ([]dto.ProjectReportDTO, error)
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
		"INSERT INTO project (id,user_id,name,created_at) VALUES ($1,$2,$3,$4)",
		project.Id, project.UserId, project.Name, project.CreatedAt,
	)
	return err
}

// GetTokenByUserId implements ITokenRepository.
func (repository *ProjectRepository) GetProjectReportByUserId(userId string) ([]dto.ProjectReportDTO, error) {
	var projects []dto.ProjectReportDTO

	row, err := repository.dbPool.Query(*repository.ctx,
		`
SELECT
    p.name AS project_name,
    e.merchant_name AS merchant_name,
    SUM(e.amount) AS total_amount,
    p.created_at AS created_at,
    COUNT(e.expense_id) AS total_expenses
FROM
    projects p
JOIN
    expenses e ON p.project_id = e.project_id
WHERE 
	p.user_id = $1
GROUP BY
    p.name, e.merchant_name, p.created_at
		`, userId,
	)

	var project dto.ProjectReportDTO
	for row.Next() {
		err = row.Scan(&project.MerchantName, &project.MerchantName, &project.TotalAmount, &project.CreatedAt, &project.TotalExpenses)
		if err != nil {
			break
		}
		projects = append(projects, project)
	}

	return projects, err
}
